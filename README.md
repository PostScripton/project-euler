# Project Euler

Personal solutions to [Project Euler](https://projecteuler.net/) problems, written mainly in Go
with some accompanying JavaScript.

## Structure

Tasks are grouped in decades of folders (`tasks_001-010`, `tasks_011-020`, `tasks_021-030`, ...).

`tasks_001-010` uses the current, expected layout — each task lives in its own zero-padded folder:

```
tasks_001-010/task_001/
├── README.md   # problem statement and link to projecteuler.net
├── go/         # solution(s), tests, and benchmark results
└── js/         # solution(s) in JavaScript
```

`tasks_011-020` and `tasks_021-030` still hold an older, ad hoc structure (standalone `.js`
files/folders, no README, no Go). These get migrated to the layout above one task at a time as
each is revisited.

## Tests

Each Go solution has a `_test.go` file with table-driven test cases. Run everything:

```shell
go test ./...
```

Run a single task:

```shell
go test ./tasks_001-010/task_001/go -v
```

A passing test means the solution produces the correct answer for the known sample/check values;
it does not by itself guarantee the code is efficient — see benchmarks below for that.

## Benchmarks

Each task's `go/benchmark.txt` holds the last recorded `go test -bench` output (ops/sec, ns/op,
memory allocations) for that solution, useful for comparing alternative approaches to the same
problem. Regenerate it for a single task with the `Taskfile.yml` at the repo root, passing the
task number:

```shell
task benchmark -- 11
```

This resolves the task number to its `tasks_XXX-YYY/task_NNN/go` folder, runs
`go test -bench=. -benchmem`, and writes the result to that task's `go/benchmark.txt`.

Lower `ns/op` and fewer `allocs/op` are better — use it to compare solutions to the same task, not
across different tasks.
