# Roadmap — sim2real Pipeline & softr Experiment Suite

> **Filed by:** strategist agent (ACMM L5 — hold-gated mode)  
> **Date:** 2026-07-29  
> **Status:** Draft — human review required before merge

This document maps the current backlog to three planning horizons. It is
intended to serve as a triage framework for the hold queue and a signal of
project trajectory for contributors and operators evaluating the sim2real
pipeline.

---

## Current State (as of 2026-07-29)

| Signal | Value |
|--------|-------|
| Governor mode | SURGE |
| Open issues (inference-sim/sim2real) | 34 |
| Hold queue | 16 issues |
| SLA violations | 18 |
| Agent PR capacity | **Blocked** — GitHub App not installed on `inference-sim` org |
| CI test coverage | 839 / 1,653 tests (50.8%) |
| Pipeline stage | Step 5 (partial); step 6 not started |

**Critical blockers before any roadmap horizon is achievable:**

1. **Install `kubestellar-hive` App on `inference-sim` org** — without this,
   agent PRs cannot be delivered to sim2real. See softr#13 and sim2real#617.

2. **Hold queue triage sprint** — 16 issues need to be classified as
   `ready / needs-design / deferred` before agent capacity can be directed
   usefully. See softr#11.

---

## Horizon 0 — Unblock (near-term, ≤2 weeks)

*Prerequisites for everything else.*

| Item | Issue | Owner |
|------|-------|-------|
| Install GitHub App on inference-sim org | sim2real#617, softr#13 | Human |
| Hold queue triage sprint (classify 16 items) | softr#11 | Human |
| Expand CI to full test suite (pytest pipeline/tests/) | quality bead 9416abfc | ci-maintainer |
| Fix `_names()` for path-string workloads (closes #572) | sim2real#572 | quality/scanner |
| Close test issues #614/#615 (scanner test noise) | sim2real#614, #615 | scanner |

---

## Horizon 1 — v0.9: Pipeline Correctness (mid-term, ≤6 weeks)

*Pipeline produces correct output for all real workload shapes.*

### GPU Capacity Probe Hardening

The capacity probe subsystem has four known correctness gaps:

| Item | Issue |
|------|-------|
| Capacity probe extractor silent on misspelled top-level keys | sim2real#271 |
| Capacity probe doesn't read `acceleratorType.labelValues` | sim2real#270 |
| Workload tolerations not located for capacity probe | sim2real#263 |
| GPU capacity probe: per-node fragmentation for multi-GPU pods | sim2real#262 |

**Milestone exit criterion:** Capacity probe produces correct GPU counts for
all workload shapes including multi-GPU pods and label-based accelerator
selection.

### Deploy Orchestration Quality

| Item | Issue |
|------|-------|
| Move interactive size probe before parallel dispatch | sim2real#207 |
| Periodic saves clobber out-of-band ConfigMap edits | sim2real#450 |
| Auto-prune stale progress-dict entries | sim2real#554 |
| Classify infra-caused PipelineRun failures, route to retry | sim2real#567 |

**Milestone exit criterion:** `deploy.py run` handles all common failure modes
(infra failures, stale state, out-of-band edits) without operator intervention.

### Documentation Parity

| Item | Issue |
|------|-------|
| Remove `cluster.py provision` from README/CLAUDE.md | sim2real#613 |
| Add Prerequisites section to README.md | sim2real#619 |
| Create CONTRIBUTING.md | sim2real#620 |
| Fix llm-d-rbac fragment documentation gap | sim2real#550 |
| Fix `generate_from_config.py` vLLM pods alias miss | sim2real#549 |

---

## Horizon 2 — v1.0: Pipeline Completeness (long-term, 8–16 weeks)

*The six-stage pipeline is complete end-to-end.*

### Calibration Infrastructure

| Item | Issue |
|------|-------|
| Per-pod calibration phase and per-pod baselining | sim2real#305 |
| Wire calibration into pipeline.yaml and deploy.py collect | sim2real#306 |

**Milestone exit criterion:** sim2real pipeline produces calibrated throughput
numbers, not just raw measurements.

### Step-6 Epic: Validate/Execute + Auto-Fix

| Item | Issue |
|------|-------|
| Epic: step-6 — Validate/execute + auto-fix | sim2real#534 |
| sim2real-check skill review pass | sim2real#487 |
| Round-trip test: translation → sim2real → validate | sim2real#592 |

**Milestone exit criterion:** A single `sim2real validate` command can
reproduce a softr-style experiment end-to-end and flag regressions.

### Security Supply-Chain Hardening

| Item | Issue |
|------|-------|
| Pin all GitHub Actions to SHAs | sim2real#626 |
| Pin Dockerfile base image to digest | sim2real#627 |
| Update requirements.txt CVE-affected jinja2/requests | sim2real#628 |
| Remove `actions/checkout@v6` (non-existent tag) | sim2real#616 |
| Pin obra/superpowers-marketplace to commit | sim2real#629 |

**Note:** All five security fixes have local branches staged by the scanner
agent. These are instant merges once GitHub App install is complete.

---

## Horizon 3 — v1.1: Adoption & Ecosystem (exploratory)

*Reduce barriers for external operators to adopt sim2real.*

- **softr experiment automation**: wire `kalantar-msb/softr` experiment bundle
  through `sim2real-bootstrap` → `sim2real` → `sim2real-check` end-to-end
  (currently softr has no `transfer.yaml` — pipeline handoff is incomplete)
- **Multi-cluster support**: sim2real across heterogeneous GPU clusters
- **Data PVC path isolation** (sim2real#553): scope `/data/` by scenario to
  prevent cross-experiment collisions
- **Warn on stale orchestrator image** (sim2real#376): surface image staleness
  in `deploy.py run --remote` before dispatch

---

## Appendix: Hold Queue Classification

This table classifies each held issue against the roadmap horizons above:

| Issue | Title (short) | Horizon | Suggested action |
|-------|--------------|---------|-----------------|
| #207 | Move size probe before dispatch | H1 | Remove hold, assign |
| #262 | GPU capacity: multi-GPU fragmentation | H1 | Remove hold, assign |
| #263 | Capacity probe workload tolerations | H1 | Remove hold, assign |
| #270 | Capacity probe acceleratorType.labelValues | H1 | Remove hold, assign |
| #271 | Capacity probe silent on misspelled keys | H1 | Remove hold, assign |
| #305 | Per-pod calibration phase | H2 | Keep hold, needs design |
| #306 | Wire calibration into pipeline | H2 | Keep hold, needs design |
| #376 | Warn stale orchestrator image | H3 | Keep hold, exploratory |
| #450 | Periodic saves clobber ConfigMap edits | H1 | Remove hold, assign |
| #487 | sim2real-check skill review pass | H2 | Keep hold, after #534 |
| #534 | Epic: step-6 validate/execute | H2 | Keep hold, needs design |
| #550 | llm-d-rbac fragment doc/code gap | H1 doc | Remove hold, assign |
| #553 | Data PVC path collision | H3 | Keep hold, exploratory |
| #554 | Auto-prune stale progress-dict entries | H1 | Remove hold, assign |
| #567 | Classify infra-caused PipelineRun failures | H1 | Remove hold, assign |
| #592 | Round-trip test | H2 | Keep hold, after step-6 |

**Recommended hold removals (Horizon 1 items):** #207, #262, #263, #270, #271,
#450, #554, #567 — these are all well-scoped, have confirmed root causes, and
agent fix branches are staged or feasible.

---

*This is a draft planning artifact filed by the strategist agent. It requires
human review and should not be treated as a committed release schedule.*
