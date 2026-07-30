package softreflectiveceiling

// Tests for the soft-reflective ceiling algorithm.
// These tests use the internal package (no external import required) to
// validate the ComputeLimit dispatch logic in isolation.

import (
	"context"
	"testing"
)

// newWrappedPolicy constructs a wrappedPolicy for testing without the
// llm-d-router external dependency.
func newWrappedPolicy(name string) *wrappedPolicy {
	return &wrappedPolicy{name: name, p: &policy{}}
}

var ctx = context.Background()

// TestComputeLimit_SingleBand verifies that a single priority band always
// gets a ceiling of 1.0 regardless of saturation.
func TestComputeLimit_SingleBand(t *testing.T) {
	w := newWrappedPolicy("test")
	for _, sat := range []float64{0.0, 0.5, 1.0} {
		ceilings := w.ComputeLimit(ctx, sat, []int{0})
		if len(ceilings) != 1 || ceilings[0] != 1.0 {
			t.Errorf("sat=%.1f: want [1.0], got %v", sat, ceilings)
		}
	}
}

// TestComputeLimit_EmptyPriorities verifies that zero bands returns an empty slice.
func TestComputeLimit_EmptyPriorities(t *testing.T) {
	w := newWrappedPolicy("test")
	ceilings := w.ComputeLimit(ctx, 0.5, nil)
	if len(ceilings) != 0 {
		t.Errorf("expected empty, got %v", ceilings)
	}
}

// TestComputeLimit_HighestPriorityAlwaysOpen verifies that band 0 is
// never gated (ceiling = 1.0) regardless of saturation.
func TestComputeLimit_HighestPriorityAlwaysOpen(t *testing.T) {
	w := newWrappedPolicy("test")
	for _, sat := range []float64{0.0, 0.5, 0.99, 1.0} {
		ceilings := w.ComputeLimit(ctx, sat, []int{0, 1, 2})
		if ceilings[0] != 1.0 {
			t.Errorf("sat=%.2f: band 0 should be 1.0, got %v", sat, ceilings[0])
		}
	}
}

// TestComputeLimit_FullSaturationBlocksLowerBands verifies that at
// saturation=1.0, all bands except band 0 are hard-blocked (ceiling=0.0).
func TestComputeLimit_FullSaturationBlocksLowerBands(t *testing.T) {
	w := newWrappedPolicy("test")
	ceilings := w.ComputeLimit(ctx, 1.0, []int{0, 1, 2})
	if ceilings[0] != 1.0 {
		t.Errorf("band 0 should be 1.0, got %v", ceilings[0])
	}
	for i := 1; i < len(ceilings); i++ {
		if ceilings[i] != 0.0 {
			t.Errorf("band %d at sat=1.0 should be 0.0, got %v", i, ceilings[i])
		}
	}
}

// TestComputeLimit_ZeroSaturationAllOpen verifies that at saturation=0.0,
// all bands are open (ceiling=1.0).
func TestComputeLimit_ZeroSaturationAllOpen(t *testing.T) {
	w := newWrappedPolicy("test")
	ceilings := w.ComputeLimit(ctx, 0.0, []int{0, 1, 2})
	for i, c := range ceilings {
		if c != 1.0 {
			t.Errorf("band %d at sat=0.0 should be 1.0, got %v", i, c)
		}
	}
}

// TestComputeLimit_ProportionalGating verifies that at mid-saturation,
// the proportional gating produces some open and some closed ticks over
// many calls (i.e., neither always-0 nor always-1).
func TestComputeLimit_ProportionalGating(t *testing.T) {
	w := newWrappedPolicy("test")
	sat := 0.5 // mid saturation — band 1 should be gated proportionally
	opens, closes := 0, 0
	for i := 0; i < 100; i++ {
		c := w.ComputeLimit(ctx, sat, []int{0, 1})
		if c[1] == 1.0 {
			opens++
		} else {
			closes++
		}
	}
	if opens == 0 {
		t.Errorf("band 1 at sat=0.5 was never open over 100 calls")
	}
	if closes == 0 {
		t.Errorf("band 1 at sat=0.5 was never closed over 100 calls")
	}
}

// TestComputeLimit_ReflectiveCeilingFormula verifies the reflective ceiling
// formula: ceiling[i] = 1 - i*sat/(N-1).
// At low saturation (sat < ceiling), all bands should be open.
func TestComputeLimit_ReflectiveCeilingBelowSat(t *testing.T) {
	w := newWrappedPolicy("test")
	// With 3 bands and sat=0.1: reflective ceilings are 1.0, 0.9, 0.8
	// Since sat=0.1 < all ceilings, all bands should be open.
	ceilings := w.ComputeLimit(ctx, 0.1, []int{0, 1, 2})
	for i, c := range ceilings {
		if c != 1.0 {
			t.Errorf("band %d at sat=0.1 should be 1.0 (below reflective ceiling), got %v", i, c)
		}
	}
}

// TestComputeLimit_BandGrowth verifies that counters grow correctly when
// a call introduces more bands than previous calls.
func TestComputeLimit_BandGrowth(t *testing.T) {
	w := newWrappedPolicy("test")
	// First call with 2 bands
	w.ComputeLimit(ctx, 0.5, []int{0, 1})
	if len(w.p.counters) < 2 {
		t.Errorf("expected at least 2 counters, got %d", len(w.p.counters))
	}
	// Second call with 4 bands — counters should grow
	w.ComputeLimit(ctx, 0.5, []int{0, 1, 2, 3})
	if len(w.p.counters) < 4 {
		t.Errorf("expected at least 4 counters after growth, got %d", len(w.p.counters))
	}
}
