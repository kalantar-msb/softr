# Roadmap — sim2real Pipeline & softr Experiment Suite

> **Filed by:** strategist agent (ACMM L5 — hold-gated mode)  
> **Date:** 2026-07-30 (cycle 2 update; original: 2026-07-29)  
> **Status:** Draft — human review required before merge

This document maps the current backlog to three planning horizons. It is
intended to serve as a triage framework for the hold queue and a signal of
project trajectory for contributors and operators evaluating the sim2real
pipeline.

---

## Current State (as of 2026-07-30, cycle 2)

> **⚠️ CRITICAL:** The hive governor only tracks `inference-sim/sim2real`. The
> `kalantar-msb/softr` repo has 33 open issues + 10 open PRs that are invisible
> to governor queue metrics. Total project debt is ~2× what SURGE mode sees.
> See issue #55.

| Signal | Value |
|--------|-------|
| Governor mode | SURGE (persistent since ≥2026-07-27) |
| Open issues (inference-sim/sim2real) | 52 |
| Open issues (kalantar-msb/softr) | 33 |
| Open issues (total, both repos) | **85** |
| Hold queue (sim2real) | 16 issues + 2 PRs |
| Hold queue (softr) | 10 PRs (all hold-gated) |
| Agent token burn rate | ~$210/hr (424M tokens in window) |
| PRs delivered to merge queue | **0** |
| Commit velocity | ~110 commits / 30 days |
| Agent PR capacity | **Blocked** — GitHub App not installed on `inference-sim` org |
| CI test coverage | ~50% (pipeline tests partially covered) |
| Pipeline stage | Step 5 complete on `main`; step 6 not started |
| Corpus-mode trace | **Landed** — sim2real#605 merged; softr automation path unblocked |

**Critical blockers before any roadmap horizon is achievable:**

1. **Install `kubestellar-hive` App on `inference-sim` org** — without this,
   agent PRs cannot be delivered to sim2real. See softr#13 and sim2real#617.

2. **Hold queue triage sprint** — 10 open softr PRs (all hold-gated) + 18 sim2real
   hold items need human review and merge/close decisions. See softr#57 and softr#11.

3. **Define SURGE exit criteria** — system is burning ~$210/hr with 0 PRs merging.
   Without exit criteria, SURGE continues indefinitely. See softr#53.

4. **Close stale ROADMAP draft PR#10** — superseded by this PR#15. See softr#28.

5. **Fix softr governor tracking** — `kalantar-msb/softr` has 33 open issues not
   counted in SURGE metrics. See softr#55.

---

## Horizon 0 — Unblock (near-term, ≤2 weeks)

*Prerequisites for everything else.*

| Item | Issue | Owner |
|------|-------|-------|
| Install GitHub App on inference-sim org | sim2real#617, softr#13 | Human |
| Define SURGE exit criteria | softr#53 | Human |
| Hold queue triage sprint (10 softr PRs, 18 sim2real items) | softr#57, softr#11 | Human |
| Fix softr repo in hive governor tracking | softr#55 | Human |
| Close stale ROADMAP PR#10 (superseded) | softr#28 | Human |
| Create CONTRIBUTING.md for softr | softr#45, softr#26 | guide *(PR#46 ready)* |
| Add Go CI workflow | softr#39, softr#49 | ci-maintainer *(PR#40 ready)* |
| Remove internal hostnames from baseline.yaml | softr#35 | scanner *(PR#37 ready)* |
| Fix data race in ComputeLimit | softr#17 | sec-check *(PR#19 ready)* |
| Expand CI to full test suite (pytest pipeline/tests/) | sim2real#632 | ci-maintainer |
| Close test issues #614/#615/#618/#624/#625 (agent test noise) | sim2real | agents |

---

## Horizon 1 — v0.9: Pipeline Correctness + softr v0.9.0 (mid-term, ≤6 weeks)

*Pipeline produces correct output for all real workload shapes. softr experiment bundle reaches v0.9.0.*

### softr v0.9.0 Exit Criteria *(new — cycle 2)*

Formal criteria for cutting the v0.9.0 release tag on `kalantar-msb/softr`. See softr#60.

| Criterion | Issue | PR Ready? |
|---|---|---|
| Data race fix merged | softr#17 | ✅ PR#19 |
| CONTRIBUTING.md merged | softr#45 | ✅ PR#46 |
| CI workflow merged | softr#39 | ✅ PR#40 |
| Internal hostnames removed | softr#35 | ✅ PR#37 |
| README matches actual repo structure | softr#30 | ✅ PR#32 |
| CHANGELOG.md with v0.9.0 entry | softr#29 | ❌ Not started |
| transfer.yaml created | softr#58 | ❌ Not started |

**Milestone exit criterion:** All 7 items above merged/closed. GitHub Release v0.9.0 cut.

### softr Workload Coverage Expansion *(new — cycle 2)*

Current BLIS results cover only 2-band (critical/sheddable) configurations. See softr#59.

| Item | Issue |
|---|---|
| Add 3-band workload YAMLs (critical/standard/sheddable) | softr#59 |
| Benchmark medium-saturation regime (0.5–0.8) | softr#59 |
| Validate N-band ceiling formula for N=3 | softr#9 |

### softr Pipeline Integration *(updated — corpus-mode landed)*

> sim2real#605 (corpus-mode trace) is **merged on main**. The softr automation path is unblocked.

| Item | Issue |
|---|---|
| Create `transfer.yaml` for softr bundle | softr#58, softr#23 |
| Use softr as step-6 demo run | softr#20, sim2real#534 |
| End-to-end: sim2real-check validates softr TTFT claims | softr#23 |

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
| Create CONTRIBUTING.md (sim2real) | sim2real#620 |
| Create CONTRIBUTING.md (softr) | softr#26 |
| Fix llm-d-rbac fragment documentation gap | sim2real#550 |
| Fix `generate_from_config.py` vLLM pods alias miss | sim2real#549 |
| Add CHANGELOG.md + version release tag | softr#29 |

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
| **File step-6 sub-issues from step-5 deferred items** | softr#20 |
| sim2real-check skill review pass | sim2real#487 |
| Round-trip test: translation → sim2real → validate | sim2real#592 |

**Gap (softr#20):** sim2real#534 has no sub-issues. The step-5 design doc has at
least 4 deferred items (cross-replica analyze aggregation, per-replica aggregate
verdict, `--replicas` shorthand on `deploy.py run`, shrink semantics) that need
to be filed before step-6 work can begin.

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

> **Note (cycle 2):** softr End-to-End Automation has been promoted to H1 (see above).
> The items below remain exploratory / long-term.

### Versioning and Release Cadence

| Item | Issue |
|------|-------|
| Add CHANGELOG.md + first version tag | softr#29 |
| Add BREAKING CHANGE annotation policy to CONTRIBUTING.md | softr#29, sim2real#620 |
| Publish GitHub Release at v0.9.0 anchor | softr#6, softr#60 |

### Other Adoption Items

- **Multi-cluster support**: sim2real across heterogeneous GPU clusters
- **Data PVC path isolation** (sim2real#553): scope `/data/` by scenario to
  prevent cross-experiment collisions
- **Warn on stale orchestrator image** (sim2real#376): surface image staleness
  in `deploy.py run --remote` before dispatch

---

## Appendix A: Hold Queue Classification (sim2real)

This table classifies each held sim2real issue against the roadmap horizons:

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
| #647 | sec-check: non-root Dockerfile USER | H1 sec | Review and merge |
| #648 | sec-check: remove cluster-wide pods RBAC | H1 sec | Review and merge |

**Recommended hold removals (Horizon 1 items):** #207, #262, #263, #270, #271,
#450, #554, #567 — well-scoped, confirmed root causes, agent fix branches staged.

---

## Appendix B: softr PR Hold Queue *(new — cycle 2)*

Ten PRs are open and hold-gated on softr. All require human review.

| PR | Branch | Type | Suggested action |
|----|--------|------|-----------------|
| #3 | guide/docs-sim2real-pipeline-status | docs | Review and merge |
| #10 | strategy/roadmap-v1 | planning (STALE) | **Close** — superseded by #15 |
| #15 | strategy/roadmap | planning | Review (this document) |
| #19 | sec/fix-data-race | security | Review and merge |
| #31 | arch/refactor-interface-consistency | refactor | Review and merge |
| #32 | arch/refactor-readme-file-structure | docs/refactor | Review and merge |
| #37 | scanner/fix-info-disclosure | security | Review and merge |
| #40 | ci/add-go-ci-workflow | CI | Review and merge |
| #43 | guide/docs-reproduction-instructions | docs | Review and merge |
| #46 | guide/docs-contributing | docs | Review and merge |

**Priority order:** PR#10 (close first), PR#37 (security), PR#19 (security),
PR#46 (CONTRIBUTING.md), PR#40 (CI), then the rest.

---

*This is a draft planning artifact filed by the strategist agent. It requires
human review and should not be treated as a committed release schedule.*
