# Configuration: Soft-Reflective Proportional Gating

## vLLM Pod Configuration

| Parameter | Value | Notes |
|-----------|-------|-------|
| Model | `Qwen/Qwen3-14B` | |
| GPU | H100-SXM-80GB | |
| `tensor_parallel_size` | 1 | |
| `max_num_seqs` | 256 | Max concurrent requests per pod |
| `max_num_batched_tokens` | 2048 | Chunked prefill budget |
| `block_size` | 16 | KV cache block size in tokens |
| `gpu_memory_utilization` | 0.9 | |
| `max_model_len` | 40960 | |
| `enable_chunked_prefill` | True | |
| Number of vLLM pods | 2 | Each gets 1 H100 GPU (TP=1) |

## llm-d EPP Configuration — Baseline

```yaml
apiVersion: llm-d.ai/v1alpha1
kind: EndpointPickerConfig
featureGates:
  - flowControl
plugins:
  - type: queue-scorer
  - type: kv-cache-utilization-scorer
  - type: prefix-cache-scorer
schedulingProfiles:
  - name: default
    plugins:
      - pluginRef: queue-scorer
        weight: 2.0
      - pluginRef: kv-cache-utilization-scorer
        weight: 2.0
      - pluginRef: prefix-cache-scorer
        weight: 3.0
flowControl: {}
```

## llm-d EPP Configuration — Treatment

```yaml
apiVersion: llm-d.ai/v1alpha1
kind: EndpointPickerConfig
featureGates:
  - flowControl
plugins:
  - type: queue-scorer
  - type: kv-cache-utilization-scorer
  - type: prefix-cache-scorer
  - type: soft-reflective-ceiling-policy
    name: soft-reflective
schedulingProfiles:
  - name: default
    plugins:
      - pluginRef: queue-scorer
        weight: 2.0
      - pluginRef: kv-cache-utilization-scorer
        weight: 2.0
      - pluginRef: prefix-cache-scorer
        weight: 3.0
flowControl:
  usageLimitPolicyPluginRef: soft-reflective
```

## llm-d EPP Configuration — Control

```yaml
apiVersion: llm-d.ai/v1alpha1
kind: EndpointPickerConfig
featureGates:
  - flowControl
plugins:
  - type: queue-scorer
  - type: kv-cache-utilization-scorer
  - type: prefix-cache-scorer
  - type: constant-ceiling-control
    name: control-policy
schedulingProfiles:
  - name: default
    plugins:
      - pluginRef: queue-scorer
        weight: 2.0
      - pluginRef: kv-cache-utilization-scorer
        weight: 2.0
      - pluginRef: prefix-cache-scorer
        weight: 3.0
flowControl:
  usageLimitPolicyPluginRef: control-policy
```

## Priority Bands

Under v0.9.0 the routerlib chart emits `InferenceObjective` resources natively from `router.inferenceObjectives`. Scenarios declare the bands as a values array; the chart generates the manifests at deploy time.

```yaml
router:
  inferenceObjectives:
    - name: critical
      priority: 100
    - name: sheddable
      priority: -50
```

## Real-Cluster Load Generator (blis observe)

```bash
blis observe \
  --server-url http://<gateway>:80 \
  --model Qwen/Qwen3-14B \
  --workload-spec <workload>.yaml \
  --max-concurrency 10000 \
  --prewarm-duration 60s \
  --warmup-requests 50 \
  --timeout 3600 \
  --post-hoc-detector composite \
  --trace-header trace.yaml \
  --trace-data trace.csv \
  --saturation-report saturation.json
```

- `--prewarm-duration 60s`: Warms CUDA kernels, EPP connections, memory allocators with small fixed requests (concurrency=4, 256 input / 64 output tokens) before real workload starts.
- `--warmup-requests 50`: Excludes the first 50 real-workload requests from the trace. Trims the brief batch-fill transient after prewarm completes.

## BLIS Simulation Flags

| Flag | Value | Notes |
|------|-------|-------|
| `--model` | qwen/qwen3-14b | |
| `--latency-model` | trained-physics | Physics-informed with learned corrections |
| `--max-model-len` | 40960 | Matches vLLM config |
| `--flow-control` | (enabled) | |
| `--saturation-detector` | utilization | max(QD/5, KV/0.8) averaged |
| `--queue-depth-threshold` | 5 | |
| `--kv-cache-util-threshold` | 0.8 | |
| `--dispatch-order` | priority | Highest priority first |
| `--num-instances` | 2 | Matches real cluster |
| `--ceiling-policy` | soft-reflective | Treatment only |
| `--usage-limit-threshold` | 1.0 | Baseline only |

## BLIS Defaults (maps to vLLM)

| BLIS default | vLLM equivalent | Value |
|---|---|---|
| max_num_seqs | --max-num-seqs | 256 |
| block_size | --block-size | 16 |
| gpu_memory_utilization | --gpu-memory-utilization | 0.9 |

## Commits

| Component | Commit/Version | Notes |
|-----------|---------------|-------|
| BLIS | a185454fc06e9c0605df04ad3d469f952d840983 | unchanged |
| llm-d-router | **v0.9.0** (plus your local patches: runner.go + softreflectiveceiling/ + constantceilingcontrol/) | was: quartic-jc branch (a6365ce1 / v0.4.0-rc.1-1645). Requires rebuild of the custom EPP image. |
| llm-d-router custom image | ghcr.io/kalantar/llm-d-router:\<new-tag\> | was: ghcr.io/kalantar/llm-d-inference-scheduler (project renamed llm-d-inference-scheduler → llm-d-router upstream — verify your registry path matches) |
| llm-d-benchmark | **v0.7.0** | was: unpinned / 5208cf38 (pre-v0.7.0). Bumping this switches chartVersions.llmDRouter → v0.9.0, images.routerEndpointPicker → v0.9.0, images.routingSidecar → v0.9.0 (repo path changed to llm-d-router-disagg-sidecar), plus scenario-side inferenceExtension.* → router.epp.* rename. |
| vLLM | **v0.23.0** (per llm-d-benchmark v0.7.0 defaults) | was: v0.20.2 — confirm whether you're pinning to v0.23.0 for the v0.9.0 run or holding at v0.20.2 |
| Nous campaign | flow-control-reflective-v2, iter-4/5 | unchanged |
