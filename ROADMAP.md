# Roadmap — sim2real Pipeline & softr Experiment Suite

> **Filed by:** strategist agent (ACMM L5 — hold-gated mode)
> **Date:** 2026-07-30 (cycle 3 update; cycle 2: 2026-07-30 09:49; original: 2026-07-29)
> **Status:** Draft — human review required before merge

This document maps the current backlog to three planning horizons. It is
intended to serve as a triage framework for the hold queue and a signal of
project trajectory for contributors and operators evaluating the sim2real
pipeline.

---

## Current State (as of 2026-07-30, cycle 3)

> **⚠️ CRITICAL (new, cycle 3):** `inference-sim/sim2real` CI is broken on main.
> The `Build Orchestrator Image` job has been failing since commit `5a0431f`
> (sim2real#661). Every agent merge lands on a broken baseline. Diagnose and fix
> before resuming SURGE mode investment in sim2real.

> **⚠️ CRITICAL (cycle 2, still open):** The hive governor tracks
> `kalantar-msb/soft-reflective` (a repo that does not exist). The actual repo
> is `kalantar-msb/softr`. Until this config is corrected, softr issues are
> invisible to the SURGE queue **and the ACMM proxy hard-blocks all writes to
> softr** (issues, PRs, comments). See softr#55.

| Signal | Value |
|--------|-------|
| Governor mode | SURGE (persistent) |
| Open issues (inference-sim/sim2real) | 49 |
| Open issues (kalantar-msb/softr) | 42 |
| Open issues (total, both repos) | **~91** |
| Open PRs (softr) | 11 (10 hold-gated + PR#63 without hold) |
| Open PRs (sim2real) | 1 (hold-gated) |
| sim2real CI status | **BROKEN on main** (sim2real#661) |
| Agent token burn rate | ~$210/hr (SURGE) |
| PRs merged to main | **0** |
| Commit velocity (sim2real) | ~40 commits / 30 days |
| GitHub App (inference-sim org) | **Not installed** — agent PRs blocked |
| ACMM proxy writes to softr | **Blocked** (softr not in governor queue) |

**Critical blockers before any roadmap horizon is achievable:**

1. **Fix sim2real CI (broken on main)** — `Build Orchestrator Image` failing
   since sec-check Dockerfile changes at `5a0431f`. No agent work can be
   validated until CI is green. See sim2real#661.

2. **Fix hive governor repo name** — change config from
   `kalantar-msb/soft-reflective` to `kalantar-msb/softr`. This unblocks the
   ACMM proxy for softr writes and adds softr to the SURGE queue. See softr#55.

3. **Install `kubestellar-hive` App on `inference-sim` org** — without this,
   agent PRs cannot be delivered to sim2real. See softr#13 and sim2real#617.

4. **Hold queue triage sprint** — 11 open softr PRs (10 hold-gated, PR#63
   without hold) + 18 sim2real hold items need human review.
   See softr#57 and softr#11.

5. **Define SURGE exit criteria** — system has been burning ~$210/hr with
   0 PRs merging. Without exit criteria, SURGE continues indefinitely.
   See softr#53.

6. **Close stale ROADMAP draft PR#10** — superseded by this PR. See softr#28.

---

## Horizon 0 — Unblock (near-term, ≤2 weeks)

*Prerequisites for everything else.*

| Item | Issue | Owner |
|------|-------|-------|
| **Fix sim2real CI (Build Orchestrator Image broken)** | sim2real#661 | Human/agent |
| Fix hive governor to track `kalantar-msb/softr` (not soft-reflective) | softr#55 | Human |
| Install GitHub App on inference-sim org | sim2real#617, softr#13 | Human |
| Define SURGE exit criteria | softr#53 | Human |
| Hold queue triage sprint (11 softr PRs, 18 sim2real items) | softr#57, softr#11 | Human |
| Close stale ROADMAP PR#10 (superseded) | softr#28 | Human |
| Create CONTRIBUTING.md for softr | softr#45, softr#26 | guide *(PR#46 ready)* |
| Add Go CI workflow | softr#39, softr#49 | ci-maintainer *(PR#40 ready)* |
| Remove internal hostnames from baseline.yaml | softr#35 | scanner *(PR#37 ready)* |
| Fix data race in ComputeLimit | softr#17 | sec-check *(PR#19 ready)* |
| Merge PR#63: add basic Go build/vet workflow | softr#39 | ci-maintainer *(no hold)* |

**Note on PR#63:** This PR was filed by the ci-maintainer agent without a
`hold` label (unlike all other agent PRs). If the intent is to allow immediate
merge, human should validate and merge. If hold-gating is required, add the
`hold` label.

---

## Horizon 1 — v0.9: Pipeline Correctness + softr v0.9.0 (mid-term, ≤6 weeks)

*Pipeline produces correct output for all real workload shapes. softr experiment bundle reaches v0.9.0.*

### softr v0.9.0 Exit Criteria

Formal criteria for cutting the v0.9.0 release tag on `kalantar-msb/softr`.
See softr#60.

| Criterion | Issue | PR Ready? |
|---|---|---|
| Data race fix merged | softr#17 | ✅ PR#19 |
| CONTRIBUTING.md merged | softr#45 | ✅ PR#46 |
| CI workflow merged | softr#39 | ✅ PR#40 or PR#63 |
| Internal hostnames removed | softr#35 | ✅ PR#37 |
| README matches actual repo structure | softr#30 | ✅ PR#32 |
| go.sum added to repo (module integrity) | softr#62 | ❌ Not started |
| CHANGELOG.md with v0.9.0 entry | softr#29 | ❌ Not started |
| transfer.yaml created | softr#58 | ❌ Not started |

**Updated:** go.sum missing added as a v0.9.0 criterion (softr#62, filed cycle 3).
**Milestone exit criterion:** All 8 items above merged/closed. GitHub Release v0.9.0 cut.

### softr Workload Coverage Expansion

Current BLIS results cover only 2-band (critical/sheddable) at high and low
saturation. See softr#59.

| Item | Issue |
|---|---|
| Add 3-band workload YAMLs (critical/standard/sheddable) | softr#59 |
| Benchmark medium-saturation regime (0.5–0.8) | softr#59 |
| Validate N-band ceiling formula for N=3 | softr#9 |
| Add decode-heavy workload variant (addresses softr#8) | softr#8 |

### softr Pipeline Integration *(corpus-mode landed)*

> sim2real#605 (corpus-mode trace) is **merged on main**. The softr automation
> path is unblocked.

| Item | Issue |
|---|---|
| Create `transfer.yaml` for softr bundle | softr#58, softr#23 |
| Use softr as step-6 demo run | softr#20, sim2real#534 |
| End-to-end: sim2real-check validates softr TTFT claims | softr#23 |

### GPU Capacity Probe Hardening

| Item | Issue |
|------|-------|
| Capacity probe extractor silent on misspelled keys | sim2real#271 |
| Capacity probe doesn't read `acceleratorType.labelValues` | sim2real#270 |
| Workload tolerations not located for capacity probe | sim2real#263 |
| GPU capacity probe: per-node fragmentation for multi-GPU pods | sim2real#262 |

### Deploy Orchestration Quality

| Item | Issue |
|------|-------|
| Move interactive size probe before parallel dispatch | sim2real#207 |
| Periodic saves clobber out-of-band ConfigMap edits | sim2real#450 |
| Auto-prune stale progress-dict entries | sim2real#554 |
| Classify infra-caused PipelineRun failures, route to retry | sim2real#567 |

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

### sim2real Architecture Debt (new, cycle 3)

The architect agent has filed a cluster of structural issues in sim2real
this cycle. These are mid-term items requiring design work before
implementation.

| Item | Issue |
|------|-------|
| sim2real.py monolith (2790 lines, 24 deferred imports) | sim2real#658 |
| REPO_ROOT computed 7× independently (fragile path resolution) | sim2real#657 |
| 58 raw subprocess calls, no kubectl abstraction layer | sim2real#656 |
| Circular dependency: assemble_run.py ↔ slicer.py | sim2real#655 |
| deploy.py monolith (3776 lines) — needs structural decomposition | sim2real#654 |

**Recommended approach:** These should be addressed as a coordinated refactor
epic rather than individual PRs, to avoid merge conflicts and overlapping
structural changes.

---

## Horizon 2 — v1.0: Pipeline Completeness (long-term, 8–16 weeks)

*The six-stage pipeline is complete end-to-end.*

### Calibration Infrastructure

| Item | Issue |
|------|-------|
| Per-pod calibration phase and per-pod baselining | sim2real#305 |
| Wire calibration into pipeline.yaml and deploy.py collect | sim2real#306 |

### Step-6 Epic: Validate/Execute + Auto-Fix

| Item | Issue |
|------|-------|
| Epic: step-6 — Validate/execute + auto-fix | sim2real#534 |
| File step-6 sub-issues from step-5 deferred items | softr#20 |
| sim2real-check skill review pass | sim2real#487 |
| Round-trip test: translation → sim2real → validate | sim2real#592 |

### sim2real Versioning & Release Cadence

sim2real has ~40 commits/30d with 0 releases and no CHANGELOG. Operators
running on live GPU clusters cannot pin to a stable version or know when
breaking changes land. See softr#29.

| Item | Issue |
|------|-------|
| Add CHANGELOG.md to sim2real | softr#29 |
| First version tag (v0.x.0) | softr#29 |
| Add BREAKING CHANGE annotation policy to CONTRIBUTING.md | softr#29 |

---

## Horizon 3 — v1.1: Adoption & Ecosystem (exploratory)

*Reduce barriers for external operators to adopt sim2real.*

### Security Supply-Chain Hardening

| Item | Issue |
|------|-------|
| Pin all GitHub Actions to SHAs | sim2real#626 |
| Pin Dockerfile base image to digest | sim2real#627 |
| Update requirements.txt CVE-affected jinja2/requests | sim2real#628 |
| Pin obra/superpowers-marketplace to commit | sim2real#629 |
| Fix RBAC ClusterRole cluster-wide pods list | sim2real#646 |
| Fix BuildKit privileged pod + unpinned images in k8s manifests | sim2real#637 |

### Other Adoption Items

- **Multi-cluster support**: sim2real across heterogeneous GPU clusters
- **Data PVC path isolation** (sim2real#553): scope `/data/` by scenario
- **Warn on stale orchestrator image** (sim2real#376): surface staleness before dispatch

---

## Appendix A: Hold Queue Classification (sim2real)

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
| #648 | sec-check: remove cluster-wide pods RBAC | H1 sec | Review and merge |

**Recommended hold removals (Horizon 1 items):** #207, #262, #263, #270, #271,
#450, #554, #567 — well-scoped, confirmed root causes, agent fix branches staged.

---

## Appendix B: softr PR Hold Queue (cycle 3 update)

11 PRs are currently open on softr. All require human review.

| PR | Branch | Type | Suggested action |
|----|--------|------|-----------------|
| #3 | guide/docs-sim2real-pipeline-status | docs | Review and merge |
| #10 | strategy/roadmap-v1 | planning (STALE) | **Close** — superseded by #15 |
| #15 | strategy/roadmap | planning | Review (cycle 2 document) |
| **This PR** | strategy/roadmap-cycle3 | planning | Review (this document, cycle 3) |
| #19 | sec/fix-data-race | security | Review and merge |
| #31 | arch/refactor-interface-consistency | refactor | Review and merge |
| #32 | arch/refactor-readme-file-structure | docs/refactor | Review and merge |
| #37 | scanner/fix-info-disclosure | security | Review and merge |
| #40 | ci/add-go-ci-workflow | CI | Review and merge (or prefer PR#63) |
| #43 | guide/docs-reproduction-instructions | docs | Review and merge |
| #46 | guide/docs-contributing | docs | Review and merge |
| #56 | ci/fix-false-green-build-vet | CI | Review and merge |
| #63 | ci/add-basic-workflow | CI (no hold) | Validate and merge or add hold |

**Priority order:** Close PR#10 and PR#15 (superseded by this PR), then:
PR#37 (security), PR#19 (security), PR#46 (CONTRIBUTING.md), PR#40 or PR#63
(CI — pick one, close the other), PR#32 (README fix), PR#56, PR#31, PR#43, PR#3.

---

*This is a draft planning artifact filed by the strategist agent. It requires
human review and should not be treated as a committed release schedule.*
