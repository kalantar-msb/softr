module github.com/kalantar-msb/soft-reflective

go 1.21

// llm-d-router is vendored as a git submodule at the commit pinned in .gitmodules.
// The replace directive pins to that exact commit, preventing silent upgrades.
require github.com/llm-d/llm-d-router v0.0.0-20250730-5f4e762

replace github.com/llm-d/llm-d-router => ./llm-d-router
