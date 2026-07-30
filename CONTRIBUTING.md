# Contributing to soft-reflective

This repository contains the `soft-reflective` proportional gating algorithm bundle for [llm-d](https://github.com/llm-d/llm-d). It was produced by [BLIS](https://github.com/inference-sim/inference-sim) (Blackbox Inference Simulator) and transferred to production via the [sim2real](https://github.com/inference-sim/sim2real) pipeline.

## Ways to contribute

- **Bug reports** — open an issue describing the observed vs expected behavior
- **Algorithm improvements** — iterate on `soft_reflective_ceiling.go` or the plugin config
- **New algorithm variants** — add a new file under `algorithms/` following the existing pattern
- **Documentation fixes** — correct or expand README.md or config.md

## Repository layout

```
algorithms/    # Go plugin source files
baselines/     # llm-d-benchmark baseline scenario YAMLs
config.md      # Full vLLM + EPP deployment configuration
results/       # BLIS simulation output (do not edit manually)
workloads/     # Workload YAML files for BLIS runs
llm-d-router/  # llm-d-router submodule (pinned to v0.9.0)
```

## Developing algorithm changes

### Prerequisites

- Go toolchain (for building + testing the plugin)
- [BLIS](https://github.com/inference-sim/inference-sim) installed and in `$PATH` (for running simulations)
- Python ≥ 3.10 (for the sim2real transfer pipeline)

### Making changes

1. Edit or add files under `algorithms/`
2. Verify the plugin compiles: `go build ./algorithms/...` (requires `llm-d-router` submodule initialized)
3. Run the BLIS simulation to validate behavior:

   ```bash
   bash scripts/run.sh        # baseline + treatment × 3 seeds
   bash scripts/compare.sh    # print comparison table
   ```

4. Update the Simulation Results table in `README.md` if results change
5. Open a pull request with your changes

### Updating the llm-d-router submodule

To test against a different version of `llm-d-router`:

```bash
cd llm-d-router
git fetch && git checkout <tag-or-sha>
cd ..
git add llm-d-router
git commit -m "chore: bump llm-d-router to <version>"
```

After bumping, re-run `scripts/run.sh` to regenerate results.

## Transferring an improved algorithm to production

If you want to transfer an improved algorithm to a real llm-d cluster, use the [sim2real pipeline](https://github.com/inference-sim/sim2real). This repository is a BLIS experiment bundle — `sim2real-bootstrap` can generate the `transfer.yaml` manifest needed for the transfer pipeline.

## Code style

- Go files: follow standard `gofmt` formatting (`gofmt -w algorithms/`)
- YAML files: 2-space indent, no trailing spaces

## Opening issues

Please search [existing issues](../../issues) before opening a new one. Include:
- The router version (`llm-d-router` submodule SHA)
- The workload(s) affected
- Observed vs expected behavior
