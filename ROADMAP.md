# Roadmap — sim2real Pipeline & softr Experiment Suite

> **Filed by:** strategist agent (ACMM L5 — hold-gated mode)  
> **Date:** 2026-08-01 (cycle 6 update; original: 2026-07-29)  
> **Status:** Draft — human review required before merge

This document maps the current backlog to three planning horizons. It is
intended to serve as a triage framework for the hold queue and a signal of
project trajectory for contributors and operators evaluating the sim2real
pipeline.

---

## Current State (as of 2026-08-01, cycle 6)

> **🚨 CRITICAL (sim2real):** The PR pipeline is in a **day-2 outage**.
> The hive PR-request watcher TLS cert issue (issue #782) remains unresolved.
> **0 PRs are open** despite 100 branches with agent work. 48+ hours with no PR
> creation. 75+ worktree-issue-* branches are orphaned.

> **🚨 CRITICAL (sim2real):** **v0.1.0 is overdue.** All prerequisites landed
> on main 2 days ago (Jul 30: ROADMAP.md, CONTRIBUTING.md, security hardening,
> CI coverage ≥97%). The only remaining step is a human cutting the release tag.

> **⚠️ CRITICAL (softr):** After 6 agent cycles, **zero PRs have been merged** into
> `kalantar-msb/softr` main. The repo has had **zero substantive code changes since
> initial commit on July 14 (18 days)**. All 13 open PRs remain hold-gated.

> **⚠️ ADVISORY MODE:** Strategist is operating at ACMM L1 (ADVISORY) — cannot
> create GitHub issues or PRs directly. All findings documented as beads only.

| Signal | Value |
|--------|-------|
| Governor mode | SURGE (persistent since ≥2026-07-27) |
| Open issues (inference-sim/sim2real) | ~46 |
| Open issues (kalantar-msb/softr) | ~34 issues + 13 PRs |
| Hold queue (sim2real) | **0 PRs** (all closed/merged — pipeline stalled) |
| Hold queue (softr) | **13 PRs** (12 hold-gated + PR#68 no-hold) |
| softr PRs merged | **0** (3 PRs closed without merge) |
| sim2real PRs merged (since Jul 30) | **20+** (ROADMAP, CONTRIBUTING, security, quality) |
| sim2real last merge | Jul 31 (PR#809) |
| Commit velocity (sim2real) | 800+ total commits, pipeline functional |
| CI test coverage (sim2real) | ~97% (quality agent delivered) |
| Pipeline stage | Step 5 complete; step 6 not started |
| sim2real v0.1.0 status | **OVERDUE** — all prereqs on main since Jul 30, no tag |
| softr v0.9.0 status | **Blocked** — 0 PRs merged, all criteria unmet |
| sim2real branches | **100 branches** (~75 orphaned worktree-issue-* branches) |

### sim2real — What Changed in Cycle 6 (vs Cycle 5)

❌ **PR pipeline stalled (day 2 outage):**
- 0 open PRs in sim2real (all previous PRs merged or closed)
- 100 branches exist: 25 guide/, 4 architect/, 3 quality/, 6 ci/, 3 scanner/, 3 sec/
- All Jul 30 agent work is stranded in branches — TLS watcher issue (#782) blocks PR creation

✅ **Already landed on main (before the stall):**
- PR#728 — ROADMAP.md (strategist)
- PR#729 — CONTRIBUTING.md + LICENSE (strategist)
- PR#809 — Security: remove dead data-pvc-explorer manifest
- Multiple security hardening PRs: SHA-pinned Actions (#691, #660, #647, #765, #795)
- RBAC: subprocess timeouts (#694), non-root Dockerfile (#647)
- Quality: 75+ new tests (#787, #801), CI test coverage ≥97%
- PR#772 — layout.repo_root() consolidation (architect)

✅ **Issues closed:** #786 (TLS delivery stall resolved), #672, #719, #805, #807

### softr — Governance Crisis

❌ **Zero merges in project lifetime.** All 3 "closed" PRs were discarded without merge:
- PR#36 (`scanner/fix-data-race`) — closed 2026-07-30, NOT merged
- PR#40 (`ci/add-go-ci-workflow`) — closed 2026-07-30, NOT merged
- PR#63 (`ci/add-basic-workflow`) — closed 2026-07-30, NOT merged

**Current main branch is identical to the July 14 initial commit.**
No CI, no CONTRIBUTING.md, no ROADMAP.md, no go.mod, no CHANGELOG.
README still references non-existent `scripts/` directory.

**Critical blockers before any softr roadmap horizon is achievable:**

1. **Hold queue triage sprint** — 13 open softr PRs need human review.
   Start with: PR#37 (security), PR#19 (data-race), PR#46 (CONTRIBUTING.md).
   See softr#57.

2. **Define SURGE exit criteria for softr** — 5 cycles with 0 merges.
   Without a review cadence, agent work has zero delivery. See softr#53.

3. **Add hold label to PR#68** — guide PR opened without required hold label.
   See softr#01645cc1 (bead).

4. **Close stale ROADMAP draft PR#10** — superseded by PR#15. See softr#28.

5. **Fix softr governor tracking** — `kalantar-msb/softr` may still be
   miscounted in SURGE metrics. See softr#55.

---

## Horizon 0 — Unblock (near-term, ≤2 weeks)

*Prerequisites for everything else.*

### sim2real (✅ largely unblocked)

| Item | Status | Issue |
|------|--------|-------|
| Install `kubestellar-hive` App on `inference-sim` org | ✅ Done | softr#13 |
| CONTRIBUTING.md + LICENSE | ✅ Merged PR#729 | sim2real#620 |
| ROADMAP.md | ✅ Merged PR#728 | — |
| TLS delivery pipeline stall | ✅ Resolved issue #786 | — |
| Security hardening pass (partial) | ✅ 5 PRs merged | sim2real#761–780 |
| Hold queue triage | ✅ 9/16 merged | sim2real#788 |
| Expand CI test suite | 🔄 In progress (97% coverage) | sim2real#632 |

### softr (🔴 fully blocked)

| Item | Status | Issue |
|------|--------|-------|
| Define SURGE exit criteria / review cadence | ❌ No action | softr#53 |
| Hold queue triage (13 softr PRs) | ❌ 0 merges | softr#57 |
| Fix softr repo in hive governor tracking | ❌ Still open | softr#55 |
| Close stale ROADMAP PR#10 | ❌ Still open | softr#28 |
| Create CONTRIBUTING.md for softr | ❌ PR#46 waiting (13+ days) | softr#45 |
| Add Go CI workflow | ❌ PR#56 waiting; PR#40 discarded | softr#39 |
| Remove internal hostnames from baseline.yaml | ❌ PR#37 waiting | softr#35 |
| Fix data race in ComputeLimit | ❌ PR#19 waiting | softr#17 |
| Add hold label to PR#68 (guide) | ❌ Policy violation | — |

---

## Horizon 1 — v0.9: Pipeline Correctness + softr v0.9.0 (mid-term, ≤6 weeks)

*Pipeline produces correct output for all real workload shapes. softr experiment bundle reaches v0.9.0.*

### sim2real v0.1.0 Release *(near-ready — cycle 5)*

All three prerequisites from cycle 4 are now met:

| Prerequisite | Status |
|---|---|
| CONTRIBUTING.md + LICENSE merged | ✅ PR#729 merged 2026-07-30 |
| ROADMAP.md merged | ✅ PR#728 merged 2026-07-30 |
| TLS delivery pipeline stall resolved | ✅ Issue #786 closed |

**Remaining before tagging v0.1.0:**
- CHANGELOG.md (not started)
- Security hardening pass partially complete — verify remaining items

### softr v0.9.0 Exit Criteria *(cycle 2 — unchanged due to zero merges)*

Formal criteria for cutting the v0.9.0 release tag on `kalantar-msb/softr`. See softr#60.

| Criterion | Issue | PR Status |
|---|---|---|
| Data race fix merged | softr#17 | ⏳ PR#19 waiting (13+ days) |
| CONTRIBUTING.md merged | softr#45 | ⏳ PR#46 waiting (13+ days) |
| CI workflow merged | softr#39 | ⏳ PR#56 waiting; PR#40 discarded |
| Internal hostnames removed | softr#35 | ⏳ PR#37 waiting (13+ days) |
| README matches actual repo structure | softr#30 | ⏳ PR#32 waiting (13+ days) |
| CHANGELOG.md with v0.9.0 entry | softr#29 | ❌ Not started |
| transfer.yaml created | softr#58 | ❌ Not started |

**Milestone exit criterion:** All 7 items above merged/closed. GitHub Release v0.9.0 cut.

> **Blocker:** All PRs are hold-gated. Zero merges in 5 cycles. No progress
> until human review sessions begin. See softr governance crisis.

### softr Workload Coverage Expansion *(cycle 2 — unchanged)*

Current BLIS results cover only 2-band (critical/sheddable) configurations. See softr#59.

| Item | Issue |
|---|---|
| Add 3-band workload YAMLs (critical/standard/sheddable) | softr#59 |
| Benchmark medium-saturation regime (0.5–0.8) | softr#59 |
| Validate N-band ceiling formula for N=3 | softr#9 |

### softr Pipeline Integration *(updated — corpus-mode still landed)*

> sim2real#605 (corpus-mode trace) is **merged on main**. The softr automation path
> remains unblocked pending transfer.yaml creation.

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

| Item | Status | Issue |
|------|--------|-------|
| Add Prerequisites section to README.md (sim2real) | ⏳ Open | sim2real#619 |
| Create CONTRIBUTING.md (sim2real) | ✅ Merged PR#729 | sim2real#620 |
| Create CONTRIBUTING.md (softr) | ⏳ PR#46 waiting | softr#26 |
| Fix llm-d-rbac fragment documentation gap | ⏳ Open | sim2real#550 |
| Add CHANGELOG.md + version release tag | ❌ Not started | softr#29 |

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

## Appendix A: Hold Queue Classification (sim2real) — *Updated Cycle 6*

**Status (2026-08-01): 0 OPEN PRs.** All previous PRs are closed or merged.
New PRs cannot be opened until the TLS watcher issue (#782) is resolved.

The following branches contain agent work that needs PRs opened:
- **guide/** (25 branches): docs updates pushed Jul 30, waiting for PR creation
- **architect/** (4 branches): refactoring work pushed Jul 30
- **quality/** (3 branches): additional test coverage
- **ci/** (6 branches): CI fixes
- **scanner/** (1 branch): RBAC fixes
- **sec/** (3 branches): security fixes

**Recommended actions (human required):**
1. Fix TLS cert issue (#782) to unblock PR watcher — OR — manually open PRs from branches
2. Cut v0.1.0 release tag on sim2real — all prereqs are on main
3. For holdfromhive issues: Remove hold from H1 items #207, #262, #263, #270, #271, #450, #554, #567

---

## Appendix B: softr PR Hold Queue — *Cycle 6 Update (0 merges, 18 days)*

13 PRs are open on softr. Zero have been merged. **18 days since initial commit with no code changes.**

| PR | Branch | Type | Status | Suggested action |
|----|--------|------|--------|-----------------|
| #3 | guide/docs-sim2real-pipeline-status | docs | ⏳ 18+ days | Review and merge |
| #10 | strategy/roadmap-v1 | planning (STALE) | ⏳ 18+ days | **Close** — superseded by #15 |
| #15 | strategy/roadmap | planning | ⏳ 18+ days | Review (this document, updated) |
| #19 | sec/fix-data-race | security | ⏳ 18+ days | Review and merge |
| #31 | arch/refactor-interface-consistency | refactor | ⏳ 18+ days | Review and merge |
| #32 | arch/refactor-readme-file-structure | docs/refactor | ⏳ 18+ days | Review and merge |
| #37 | scanner/fix-info-disclosure | security | ⏳ 18+ days | Review and merge |
| #43 | guide/docs-reproduction-instructions | docs | ⏳ 18+ days | Review and merge |
| #46 | guide/docs-contributing | docs | ⏳ 18+ days | Review and merge |
| #56 | ci/fix-false-green-build-vet | CI | ⏳ 18+ days | Review and merge |
| #65 | sec/fix-go-sum | security | ⏳ 18+ days | Review and merge |
| #66 | ci/add-algorithm-tests | CI | ⏳ 18+ days | Review and merge |
| #68 | guide/docs-results-directory | docs | ❌ **NO HOLD LABEL** | Add hold label FIRST; then review |

**Discarded (closed without merge):** PR#36 (data-race), PR#40 (CI workflow), PR#63 (CI workflow)

**Priority order (cycle 6):**
1. **PR#68** — Add `hold` label immediately (governance: agent PR without hold label)
2. **PR#37, PR#19** — Security fixes (data race, info disclosure)
3. **PR#10** — Close (stale, superseded by #15)
4. **PR#65** — go.sum integrity
5. **PR#56, PR#66** — CI fixes (build/vet, algorithm tests)
6. **PR#46** — CONTRIBUTING.md
7. Remaining docs PRs (#3, #32, #43)

---

*This is a draft planning artifact filed by the strategist agent. It requires
human review and should not be treated as a committed release schedule.*
