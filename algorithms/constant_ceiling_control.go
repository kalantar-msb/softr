package constantceilingcontrol

import (
	"context"
	"encoding/json"

	"github.com/llm-d/llm-d-router/pkg/epp/framework/interface/plugin"
	"github.com/llm-d/llm-d-router/pkg/epp/framework/plugins/flowcontrol/usagelimits"
)

const PolicyType = "constant-ceiling-control"

// Factory creates a constant ceiling control policy (always returns 1.0 for all bands).
// This is the control variant — functionally identical to baseline (no gating).
func Factory(name string, _ *json.Decoder, _ plugin.Handle) (plugin.Plugin, error) {
	return usagelimits.NewPolicyFunc(name, computeConstantCeilings), nil
}

func computeConstantCeilings(_ context.Context, _ float64, priorities []int) []float64 {
	ceilings := make([]float64, len(priorities))
	for i := range ceilings {
		ceilings[i] = 1.0
	}
	return ceilings
}
