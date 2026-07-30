package constantceilingcontrol

// Tests for the constant-ceiling-control policy.
// This package uses usagelimits.NewPolicyFunc — tested via a simple
// integration-style check since the logic is a trivial constant function.

import (
	"context"
	"testing"
)

var ctx = context.Background()

// computeConstantCeilings is the pure function we can test directly
// without any external dependency.

// TestComputeConstantCeilings_ReturnsAllOnes verifies that all priorities
// receive a ceiling of 1.0.
func TestComputeConstantCeilings_ReturnsAllOnes(t *testing.T) {
	ceilings := computeConstantCeilings(ctx, 0.5, []int{0, 1, 2})
	for i, c := range ceilings {
		if c != 1.0 {
			t.Errorf("priority %d: expected 1.0, got %v", i, c)
		}
	}
}

// TestComputeConstantCeilings_EmptyPriorities returns empty slice.
func TestComputeConstantCeilings_EmptyPriorities(t *testing.T) {
	ceilings := computeConstantCeilings(ctx, 1.0, nil)
	if len(ceilings) != 0 {
		t.Errorf("expected empty slice, got %v", ceilings)
	}
}

// TestComputeConstantCeilings_HighSaturationStillAllOnes verifies that
// even at saturation=1.0, ceilings are all 1.0 (control/no gating).
func TestComputeConstantCeilings_HighSaturationStillAllOnes(t *testing.T) {
	ceilings := computeConstantCeilings(ctx, 1.0, []int{0, 1, 2, 3})
	for i, c := range ceilings {
		if c != 1.0 {
			t.Errorf("priority %d at sat=1.0: expected 1.0, got %v", i, c)
		}
	}
}

// TestComputeConstantCeilings_LengthMatchesPriorities verifies output
// length matches input length.
func TestComputeConstantCeilings_LengthMatchesPriorities(t *testing.T) {
	for _, n := range []int{1, 3, 5, 10} {
		priorities := make([]int, n)
		ceilings := computeConstantCeilings(ctx, 0.5, priorities)
		if len(ceilings) != n {
			t.Errorf("n=%d: expected %d ceilings, got %d", n, n, len(ceilings))
		}
	}
}
