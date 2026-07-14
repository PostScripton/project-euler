Add or continue a solution for a Project Euler task in this repository.

Usage: `/add-solution <task number>` (e.g. `/add-solution 14`)

## Step 1 — Locate the task

Figure out which range folder the task belongs to (`tasks_001-010`, `tasks_011-020`,
`tasks_021-030`, ...) and check what already exists for it:

- **New structure** (already correct): `tasks_0XX-0YY/task_0NN/` with `README.md`, `go/`, `js/`.
  Nothing to migrate — go straight to Step 3.
- **Old structure** (needs migration first): a standalone file/folder directly under the range
  folder, e.g. `tasks_011-020/task_11.js` or `tasks_011-020/task_13/task.js` — no `README.md`,
  no `go/`. This is legacy and must be migrated before adding anything new (see Step 2).
- **Doesn't exist yet**: no file/folder for this task number anywhere. Create
  `tasks_0XX-0YY/task_0NN/` and fetch `README.md` from `https://projecteuler.net/problem=N` as
  described in Step 2.2 below, then go to Step 3.

## Step 2 — Migrate old structure to new structure (only if old structure found)

1. Create `tasks_0XX-0YY/task_0NN/` (zero-padded, matching the `tasks_001-010/task_001` pattern).
2. Fetch the problem from `https://projecteuler.net/problem=N` and write `README.md` from the
   live page content — never reuse or trust the comment header embedded in the old `.js` file
   (it may be a paraphrase, a translation, or stale). Format it like
   `tasks_001-010/task_001/README.md` — title, `**Link:** https://projecteuler.net/problem=N`,
   `## Definition` with the statement converted to markdown.
3. Move any supplementary files the problem ships with (large data files, e.g.
   `tasks_011-020/task_13/number.txt`) to sit directly in `tasks_0XX-0YY/task_0NN/`, next to
   `README.md` — not inside `go/` or `js/`, since they're shared across languages.
4. Move the existing JS code into `js/solution_1.js`, stripped of the problem-statement comment
   header and any stray `console.log` demonstration lines — keep just the solution function(s).
5. Delete the old file/folder once the contents are ported over.
6. Confirm the migration with the user before moving on to writing the new solution — this is a
   structural change to existing work, not just an addition.

## Step 3 — Add the new solution

Follow the pattern in `tasks_001-010` for whichever language(s) the user wants:

- **Go**: `go/solution_N.go` (package `solution`, function `FirstSolution`/`SecondSolution`/...),
  `go/solution_N_test.go` (table-driven over the shared `tests`, `t.Parallel()`, a
  `BenchmarkFirstSolution`-style bench), and `go/solutions_test.go` if it doesn't exist yet
  (`tests` slice + `benchmarkNumber` constant, `args` struct matching the function's params).
- **JS**: `js/solution_N.js`.

If the problem provides supplementary data (a large input file, a word/name list, etc.), it goes
directly in `tasks_0XX-0YY/task_0NN/` next to `README.md`, not duplicated inside `go/`/`js/` —
both languages' solutions read it from there.

Boilerplate (test scaffolding, file layout) can be written directly. The actual algorithm is a
teaching exercise: follow this repo's tutoring rules — guide with questions and hints, don't hand
over a complete implementation, unless the user has already solved it and explicitly asks for a
reference solution or alternative approach.

## Step 4 — Verify

```shell
go test ./tasks_0XX-0YY/task_0NN/go -v
```
