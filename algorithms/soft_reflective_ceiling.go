package softreflectiveceiling

import (
	"context"
	"encoding/json"
	"math"
	"sync"
	"sync/atomic"

	"github.com/llm-d/llm-d-router/pkg/epp/framework/interface/plugin"
	"github.com/llm-d/llm-d-router/pkg/epp/framework/plugins/flowcontrol/usagelimits"
)

const PolicyType = "soft-reflective-ceiling-policy"

// bandState holds per-band tick counter and the last-used period.
// Resetting the counter on period change ensures the proportional gating
// ratio is accurate immediately after a saturation transition.
type bandState struct {
	counter    atomic.Int64
	lastPeriod atomic.Int64
}

// policy holds per-band state for proportional dispatch gating.
// mu protects bands slice growth; individual atomic fields within
// bandState are lock-free for the hot path.
type policy struct {
	mu    sync.Mutex
	bands []bandState
}

// Factory creates a soft-reflective ceiling policy plugin instance.
// Uses the framework-provided NewPolicyFunc helper for consistent TypedName
// behavior, matching the pattern used by the control plugin and upstream policies.
func Factory(name string, _ *json.Decoder, _ plugin.Handle) (plugin.Plugin, error) {
	p := &policy{}
	return usagelimits.NewPolicyFunc(name, p.computeLimit), nil
}

// computeLimit implements the soft-reflective proportional gating algorithm.
//
// Ceiling formula: ceiling[i] = 1 - i*saturation/(N-1)
// Band 0 (highest priority): ceiling = 1.0 always (never gated)
// Band i (i>0): proportional dispatch when saturation >= ceiling[i]
func (p *policy) computeLimit(_ context.Context, saturation float64, priorities []int) []float64 {
	n := len(priorities)
	ceilings := make([]float64, n)

	if n <= 1 {
		if n == 1 {
			ceilings[0] = 1.0
		}
		return ceilings
	}

	// Grow bands slice if new priority bands appear (protected by mutex).
	if len(p.bands) < n {
		p.mu.Lock()
		for len(p.bands) < n {
			p.bands = append(p.bands, bandState{})
		}
		p.mu.Unlock()
	}

	for i := range priorities {
		if i == 0 {
			ceilings[i] = 1.0 // highest priority: never gated
			continue
		}

		// Reflective ceiling: ceiling[i] = 1 - i*sat/(N-1)
		reflectiveCeiling := 1.0 - float64(i)*saturation/float64(n-1)

		if saturation < reflectiveCeiling {
			// Below ceiling: dispatch freely
			ceilings[i] = 1.0
		} else if saturation >= 1.0 {
			// Fully saturated: hard block
			ceilings[i] = 0.0
		} else {
			// Proportional gating: dispatch every period-th tick
			period := int64(math.Max(1, math.Round(saturation/(1.0-saturation+1e-9))))

			// Reset counter when period changes to ensure immediate accuracy
			// of the proportional ratio after saturation transitions.
			if prev := p.bands[i].lastPeriod.Swap(period); prev != period {
				p.bands[i].counter.Store(0)
			}

			tick := p.bands[i].counter.Add(1)
			if tick%period == 0 {
				ceilings[i] = 1.0 // open gate this tick
			} else {
				ceilings[i] = 0.0 // closed this tick
			}
		}
	}

	return ceilings
}
