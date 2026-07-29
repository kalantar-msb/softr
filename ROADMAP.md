# ROADMAP: Soft-Reflective Proportional Gating

> **Status:** Active development — first experiment bundle generated July 2026
>
> This roadmap captures near-, mid-, and long-term priorities for the
> `soft-reflective-ceiling-policy` llm-d plugin and the sim2real workflow
> that produced it.

## Near-Term (Q3 2026)

### 1. v0.1.0 Release [![issue](https://img.shields.io/badge/issue-%236-blue)](https://github.com/kalantar-msb/softr/issues/6)

Tag the current state as `v0.1.0` with a GitHub Release. Establish the compatibility matrix:

| softr | llm-d-router | Go |
|---|---|---|
| v0.1.0 | v0.9.0 (`5f4e762f`) | 1.21+ |

The release body should include the BLIS simulation results table from README.md.

### 2. CI Pipeline [![issue](https://img.shields.io/badge/issue-%237-blue)](https://github.com/kalantar-msb/softr/issues/7)

Add `.github/workflows/ci.yml` with:
- `go vet ./...` on `algorithms/`
- `go test ./...` unit tests for `ComputeLimit()` edge cases
- Build validation against the pinned llm-d-router submodule

Baseline coverage target: 80% on `soft_reflective_ceiling.go`.

### 3. Pipeline Status Documentation

Resolve outstanding documentation gap (see issues #3, #4): add `Sim2Real Pipeline Status`
section to README.md documenting transfer.yaml location, bootstrap status, and baselines/ hand-correction.

---

## Mid-Term (Q4 2026)

### 4. Reasoning Workload Algorithm Tuning [![issue](https://img.shields.io/badge/issue-%238-blue)](https://github.com/kalantar-msb/softr/issues/8)

The current algorithm is not effective for decode-heavy reasoning workloads (KV saturation
permanently at ≥ 1.0). Research alternative saturation signals:

- Queue depth per priority band
- TTFOT (time-to-first-output-token) as a per-class proxy
- Preemption rate as a decode pressure signal

Target: achieve > 10% TTFT improvement for reasoning workloads (currently -2.9%).

### 5. N-Band Validation [![issue](https://img.shields.io/badge/issue-%239-blue)](https://github.com/kalantar-msb/softr/issues/9)

Validate the algorithm for N=3 priority bands (critical / batch / background).
Add a 3-band workload to `workloads/` and run BLIS simulation.
Add config.md guidance for multi-class `InferenceObjective` CRD deployment.

### 6. CONTRIBUTING.md + Contributor Onboarding

Currently 100% single-contributor. Add:
- `CONTRIBUTING.md` with development setup, test instructions, and PR conventions
- Good-first-issue labels on clearly scoped tasks
- A developer guide in `docs/` covering the BLIS simulation workflow

---

## Long-Term (2027)

### 7. Integration with llm-d Upstream

Propose `soft-reflective-ceiling-policy` for inclusion in the official
`llm-d-router` plugin registry. This requires:
- Upstream interface stability (track llm-d `flowcontrol` API changes)
- E2E integration tests on real hardware (not just BLIS simulation)
- A formal performance report comparing against baseline and static ceiling

### 8. Adaptive Saturation Calibration

The current policy uses a fixed saturation signal. Explore adaptive calibration:
- Auto-tune the ceiling formula coefficients based on observed TTFT distributions
- Integrate with llm-d's SaturationDetector for per-workload feedback loops

---

## Issue Index

| # | Title | Horizon | Status |
|---|---|---|---|
| [#6](https://github.com/kalantar-msb/softr/issues/6) | No versioned release | near-term | open |
| [#7](https://github.com/kalantar-msb/softr/issues/7) | No CI pipeline | near-term | open |
| [#8](https://github.com/kalantar-msb/softr/issues/8) | Reasoning workload tuning | mid-term | open |
| [#9](https://github.com/kalantar-msb/softr/issues/9) | N-band validation | mid-term | open |

---

*Maintained by the strategist agent. Last updated: 2026-07-29.*
