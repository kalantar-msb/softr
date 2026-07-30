# Sim2Real Bundle: Soft-Reflective Proportional Gating

A `UsageLimitPolicy` plugin for llm-d that reduces critical-class TTFT by 40-91% at <3% throughput cost using proportional dispatch gating.

## File Structure

```
soft-reflective/
├── README.md
├── config.md                           # Full deployment configs
├── algorithms/
│   ├── soft_reflective_ceiling.go      # Treatment plugin
│   └── constant_ceiling_control.go     # Control plugin (ceiling=1.0)
├── scripts/
│   ├── run.sh                          # BLIS simulation (baseline + treatment × 3 seeds)
│   ├── compare.sh                      # Results comparison table
│   └── treatment.patch                 # BLIS patch for simulation
└── workloads/
    ├── interactive_chat.yaml           # 150 QPS, 30/70, short tokens
    ├── code_generation.yaml            # 18 QPS, 30/70, large input
    └── reasoning.yaml                  # 1 QPS, 50/50, huge output
```

## Algorithm

**Ceiling formula:**
```
ceiling[i] = 1 - i * saturation / (N-1)
```
- Band 0 (critical): ceiling = 1.0 always
- Band 1 (sheddable): ceiling = 1 - saturation

**Proportional enforcement** (the key difference from binary policies):

When `saturation >= ceiling` for a band, instead of fully blocking:
```
period = round(saturation / (1 - saturation))
dispatch every period-th tick (via internal counter)
```

This is encoded entirely within the `ComputeLimit()` return values — the plugin alternates between returning ceiling=1.0 (open) and ceiling=0.0 (block) on successive dispatch ticks. No core llm-d code changes required.

## llm-d Migration

### Step 1: Add plugin file

Place `algorithms/soft_reflective_ceiling.go` at:
```
pkg/epp/framework/plugins/flowcontrol/usagelimits/softreflectiveceiling/policy.go
```

The plugin implements `flowcontrol.UsageLimitPolicy` interface defined at:
- **Interface:** `pkg/epp/framework/interface/flowcontrol/plugins.go:159-178`
- **Method:** `ComputeLimit(ctx context.Context, saturation float64, priorities []int) (ceilings []float64)`
- **Helper:** `pkg/epp/framework/plugins/flowcontrol/usagelimits/usagelimitpolicy.go` (NewConstPolicy at line 60, NewPolicyFunc at line 71)

### Step 2: Register in runner.go

At `cmd/epp/runner/runner.go:533` (after the existing `StaticUsageLimitPolicyType` registration), add:

```go
fwkplugin.Register(softreflectiveceiling.PolicyType, softreflectiveceiling.Factory)
```

**Existing pattern** (line 533):
```go
fwkplugin.Register(usagelimits.StaticUsageLimitPolicyType, usagelimits.StaticPolicyFactory)
```

### Step 3: Import

Add to imports in `cmd/epp/runner/runner.go`:
```go
"github.com/llm-d/llm-d-router/pkg/epp/framework/plugins/flowcontrol/usagelimits/softreflectiveceiling"
```

### Step 4: Build

```bash
go build ./cmd/epp/...
```

No new dependencies. No go.mod changes.

## How the Dispatch Cycle Consumes This Plugin

The dispatch cycle in `pkg/epp/flowcontrol/controller/internal/processor.go:342-395` (function `dispatchCycle`):

```go
ceilings := sp.usageLimitPolicy.ComputeLimit(ctx, saturation, priorities)  // line 355

for i, priority := range priorities {
    usageLimit := ceilings[i]
    if saturation >= usageLimit {  // line 361
        return false  // HoL block (block extends through line 367)
    }
    // ... select and dispatch item
}
```

The plugin controls dispatch by returning:
- `ceiling = 1.0` → `saturation < 1.0` → dispatch proceeds
- `ceiling = 0.0` → `saturation >= 0.0` → HoL block fires

By alternating these values across ticks using internal counters, the plugin creates proportional throughput without modifying the dispatch cycle code.

## Signal Mapping

| Signal | llm-d source |
|--------|-------------|
| Saturation | `SaturationDetector.Saturation(ctx, pool)` — passed as arg to `ComputeLimit` |
| Priorities | `registry.AllOrderedPriorityLevels()` — passed as arg to `ComputeLimit` |
| Dispatch tick | Every 1ms (`processor.go:182`: `sp.clock.NewTicker(time.Millisecond)`) |
| Priority bands | Dynamic — created from `InferenceObjective` CRs |

## Deployment (3 variants)

| Variant | Plugin | Config |
|---------|--------|--------|
| Baseline | none | `flowControl: {}` |
| Control | `constant-ceiling-control` | `usageLimitPolicyPluginRef: control-policy` |
| Treatment | `soft-reflective-ceiling-policy` | `usageLimitPolicyPluginRef: soft-reflective` |

The ceiling policy is the only variable. See `config.md` for full YAML configs.

<!-- Note (v0.9.0): the Config column shows the fields as they appear inside
     the EndpointPickerConfig YAML embedded in router.epp.pluginsCustomConfig.
     The wrapping key changed from inferenceExtension.pluginsCustomConfig to
     router.epp.pluginsCustomConfig in llm-d-benchmark v0.7.0. -->

## Simulation Results (BLIS)

| Workload | Rate | Split | Crit TTFT mean Δ | Crit TTFT P90 Δ | Throughput cost |
|---|---|---|---|---|---|
| interactive_chat | 150 QPS | 30/70 | **-40.7%** | -29.7% | -0.4% |
| code_generation | 18 QPS | 30/70 | **-91.3%** | -89.6% | -2.5% |
| reasoning | 1 QPS | 50/50 | -2.9% | -0.2% | +2.0% |

Best on prefill-heavy workloads with meaningful sheddable fraction (≥50%). Minimal effect on decode-heavy workloads where KV saturation pins at ≥1.0 permanently.

### Raw results: `results/`

The `results/` directory contains the raw BLIS simulation outputs used to produce the table above.

**File naming:** `{variant}_{workload}_{seed}.{ext}`

| Segment | Values | Meaning |
|---------|--------|---------|
| `variant` | `baseline`, `treatment` | Policy under test: no ceiling (baseline) or soft-reflective proportional gating (treatment) |
| `workload` | `interactive_chat`, `code_generation`, `reasoning` | Workload type (see `workloads/` for the YAML definitions) |
| `seed` | `s42`, `s43`, `s44` | Three independent seeds per workload; delta figures in the table above are averaged across all three |
| `ext` | `.json` | Per-request metrics array with aggregate statistics (`completed_requests`, `responses_per_sec`, `ttft_mean_ms`, `ttft_p90_ms`, `itl_mean_ms`, etc.) |
| `ext` | `.txt` | Human-readable simulation log with BLIS startup warnings and the `=== Simulation Metrics ===` block |

The `.json` `instance_id: "cluster"` aggregate carries the cluster-level statistics used to compute the table above (averaging the three seeds per workload). The `workloads/` directory contains the per-rate YAML inputs that produced these results. Per-instance breakdowns (`instance_id: "instance_0"`) appear in the `.txt` file.

## Code References

Verified against llm-d-router v0.9.0 (submodule commit `5f4e762f`).

| What | File | Lines |
|------|------|-------|
| UsageLimitPolicy interface | `pkg/epp/framework/interface/flowcontrol/plugins.go` | 159-178 |
| ComputeLimit called | `pkg/epp/flowcontrol/controller/internal/processor.go` | 355 |
| Saturation check (HoL block) | `pkg/epp/flowcontrol/controller/internal/processor.go` | 361-367 |
| Dispatch tick interval | `pkg/epp/flowcontrol/controller/internal/processor.go` | 182 |
| Existing static policy | `pkg/epp/framework/plugins/flowcontrol/usagelimits/usagelimitpolicy.go` | 60-77 (NewConstPolicy 60-68, NewPolicyFunc 71-77) |
| Plugin registration | `cmd/epp/runner/runner.go` | 533 |
| SaturationDetector interface | `pkg/epp/framework/interface/flowcontrol/plugins.go` | 124-137 |
| dispatchCycle function | `pkg/epp/flowcontrol/controller/internal/processor.go` | 342-395 |
