package softreflectiveceiling

import (
	"context"
	"encoding/json"
	"math"
	"sync"
	"sync/atomic"

	"github.com/llm-d/llm-d-router/pkg/epp/framework/interface/plugin"
)

const PolicyType = "soft-reflective-ceiling-policy"

// policy holds per-band tick counters for proportional dispatch.
// mu protects the counters slice during growth; individual counter
// operations use atomic.Int64 and do not need the mutex.
type policy struct {
	mu       sync.Mutex
	counters []atomic.Int64
}

// Factory creates a soft-reflective ceiling policy plugin instance.
func Factory(name string, _ *json.Decoder, _ plugin.Handle) (plugin.Plugin, error) {
	p := &policy{}
	return &wrappedPolicy{name: name, p: p}, nil
}

type wrappedPolicy struct {
	name string
	p    *policy
}

func (w *wrappedPolicy) TypedName() plugin.TypedName {
	return plugin.TypedName{Type: PolicyType + "-type", Name: w.name}
}

func (w *wrappedPolicy) ComputeLimit(ctx context.Context, saturation float64, priorities []int) []float64 {
	n := len(priorities)
	ceilings := make([]float64, n)

	if n <= 1 {
		if n == 1 {
			ceilings[0] = 1.0
		}
		return ceilings
	}

	// Grow counters if new bands appear.
	// The mutex protects the slice header (len/cap/ptr) during append;
	// subsequent atomic operations on stable elements are lock-free.
	w.p.mu.Lock()
	for len(w.p.counters) < n {
		w.p.counters = append(w.p.counters, atomic.Int64{})
	}
	w.p.mu.Unlock()

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
			tick := w.p.counters[i].Add(1)
			if tick%period == 0 {
				ceilings[i] = 1.0 // open gate this tick
			} else {
				ceilings[i] = 0.0 // closed this tick
			}
		}
	}

	return ceilings
}
