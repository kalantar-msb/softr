//go:build ignore
// +build ignore

// Standalone validation of soft-reflective ceiling algorithm.
// Run with: go run ./ceiling_test.go
// This test validates the mathematical properties of ComputeLimit
// without requiring the llm-d-router module.

package main

import (
	"fmt"
	"math"
	"os"
	"sync/atomic"
)

// Mirror of the production policy struct (copied to avoid import)
type policy struct {
	counters []atomic.Int64
}

// computeLimit mirrors soft_reflective_ceiling.go ComputeLimit
func computeLimit(p *policy, saturation float64, priorities []int) []float64 {
	n := len(priorities)
	ceilings := make([]float64, n)
	if n <= 1 {
		if n == 1 {
			ceilings[0] = 1.0
		}
		return ceilings
	}
	for len(p.counters) < n {
		p.counters = append(p.counters, atomic.Int64{})
	}
	for i := range priorities {
		if i == 0 {
			ceilings[i] = 1.0
			continue
		}
		reflectiveCeiling := 1.0 - float64(i)*saturation/float64(n-1)
		if saturation < reflectiveCeiling {
			ceilings[i] = 1.0
		} else if saturation >= 1.0 {
			ceilings[i] = 0.0
		} else {
			period := int64(math.Max(1, math.Round(saturation/(1.0-saturation+1e-9))))
			tick := p.counters[i].Add(1)
			if tick%period == 0 {
				ceilings[i] = 1.0
			} else {
				ceilings[i] = 0.0
			}
		}
	}
	return ceilings
}

type testCase struct {
	name       string
	saturation float64
	priorities []int
	checkFn    func(ceilings []float64) error
}

func main() {
	p := &policy{}
	failures := 0

	cases := []testCase{
		{
			name:       "n=1 always gets ceiling 1.0",
			saturation: 0.8,
			priorities: []int{0},
			checkFn: func(c []float64) error {
				if c[0] != 1.0 {
					return fmt.Errorf("expected 1.0, got %v", c[0])
				}
				return nil
			},
		},
		{
			name:       "n=0 returns empty slice",
			saturation: 0.5,
			priorities: []int{},
			checkFn: func(c []float64) error {
				if len(c) != 0 {
					return fmt.Errorf("expected empty, got %v", c)
				}
				return nil
			},
		},
		{
			name:       "band 0 always gets ceiling 1.0 (high sat)",
			saturation: 0.9,
			priorities: []int{0, 1},
			checkFn: func(c []float64) error {
				if c[0] != 1.0 {
					return fmt.Errorf("band 0 ceiling should be 1.0, got %v", c[0])
				}
				return nil
			},
		},
		{
			name:       "sat=0: all bands get ceiling 1.0",
			saturation: 0.0,
			priorities: []int{0, 1, 2},
			checkFn: func(c []float64) error {
				for i, v := range c {
					if v != 1.0 {
						return fmt.Errorf("band %d should be 1.0 at sat=0, got %v", i, v)
					}
				}
				return nil
			},
		},
		{
			name:       "sat=1.0: only band 0 gets ceiling 1.0",
			saturation: 1.0,
			priorities: []int{0, 1, 2},
			checkFn: func(c []float64) error {
				if c[0] != 1.0 {
					return fmt.Errorf("band 0 should be 1.0, got %v", c[0])
				}
				for i := 1; i < len(c); i++ {
					if c[i] != 0.0 {
						return fmt.Errorf("band %d should be 0.0 at sat=1.0, got %v", i, c[i])
					}
				}
				return nil
			},
		},
		{
			name:       "ceiling count: n=2 bands",
			saturation: 0.5,
			priorities: []int{0, 1},
			checkFn: func(c []float64) error {
				if len(c) != 2 {
					return fmt.Errorf("expected 2 ceilings, got %d", len(c))
				}
				return nil
			},
		},
		{
			name:       "ceiling count: n=3 bands",
			saturation: 0.5,
			priorities: []int{0, 1, 2},
			checkFn: func(c []float64) error {
				if len(c) != 3 {
					return fmt.Errorf("expected 3 ceilings, got %d", len(c))
				}
				return nil
			},
		},
		{
			name:       "ceilings are in [0,1]",
			saturation: 0.5,
			priorities: []int{0, 1, 2, 3},
			checkFn: func(c []float64) error {
				for i, v := range c {
					if v < 0.0 || v > 1.0 {
						return fmt.Errorf("band %d ceiling %v out of [0,1]", i, v)
					}
				}
				return nil
			},
		},
	}

	for _, tc := range cases {
		ceilings := computeLimit(p, tc.saturation, tc.priorities)
		if err := tc.checkFn(ceilings); err != nil {
			fmt.Printf("FAIL %s: %v\n", tc.name, err)
			failures++
		} else {
			fmt.Printf("PASS %s\n", tc.name)
		}
	}

	if failures > 0 {
		fmt.Printf("\n%d test(s) failed\n", failures)
		os.Exit(1)
	}
	fmt.Printf("\nAll %d tests passed\n", len(cases))
}
