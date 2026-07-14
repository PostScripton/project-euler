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
2. Fetch the problem statement from `https://projecteuler.net/problem=N`, in this order:
   1. **Try the Chrome browser tool first** (`mcp__claude-in-chrome__*`, via `navigate` +
      `get_page_text`). This is the preferred source since it renders the real page.
   2. If the extension isn't connected, or the direct page load hits Project Euler's bot-check
      (a "Security Verification" / reCAPTCHA page — this also happens with a plain `curl`/`WebFetch`
      GET), fall back to `WebSearch` for the problem's title and statement text, cross-checking
      against a couple of the returned mirror/solution pages.
   3. Whichever path works, write `README.md` from that content — never reuse or trust the
      comment header embedded in the old `.js` file (it may be a paraphrase, a translation, or
      stale). Format it like `tasks_001-010/task_001/README.md` — title,
      `**Link:** https://projecteuler.net/problem=N`, `## Definition` with the statement
      converted to markdown.
   4. Tell the user which of the two sources was actually used (direct page vs. web search),
      since the web-search fallback is reconstructed from third-party mirrors rather than fetched
      straight from the source.

   **Precedent (task 11):** the Chrome extension wasn't connected, and a direct `curl`/`WebFetch`
   GET to `https://projecteuler.net/problem=11` returned Project Euler's "Security Verification"
   reCAPTCHA page — so the fallback was used. Query: `WebSearch("\"Largest product in a grid\"
   project euler problem 11 full statement")`. It surfaced the official page
   `https://projecteuler.net/problem=11` alongside mirrors (`euler.stephan-brumme.com/11/`,
   `betaprojects.com/solutions/project-euler/project-euler-problem-011-solution/`,
   `euler.beerbaronbill.com/en/latest/solutions/11.html`, `raw.org/puzzle/project-euler/problem-11/`,
   `atyansh.com/euler/11/`), and the search tool's own synthesized summary of those results supplied
   the statement text used in the README.
3. Move any sizeable data the problem provides — a downloadable file, a grid, a list, a large
   number — into its own file (e.g. `grid.txt`, `number.txt`) sitting directly in
   `tasks_0XX-0YY/task_0NN/`, next to `README.md`, not inside `go/` or `js/`. This applies even
   when the data is only given inline in the problem statement (like task 11's 20×20 grid) —
   extract it out rather than hardcoding it as a literal in each language's source, since it's
   shared across languages and shouldn't be duplicated.
4. Move the existing JS code into `js/solution_1.js`, stripped of the problem-statement comment
   header — keep just the solution function(s) (including any demonstration `console.log`, to
   match the existing `task_001` style). Update it to read the extracted data file instead of
   using an inline literal.
5. Delete the old file/folder once the contents are ported over.
6. Confirm the migration with the user before moving on to writing the new solution — this is a
   structural change to existing work, not just an addition.
7. Commit in two steps (see "Commits" below): first the new-structure scaffolding (README + data
   file), then the ported legacy solution.

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

## Commits

Keep this to about **3 commits per task** — not one atomic commit per file. Claude is responsible
for the first two; the user adds the third themselves once they've written the actual solution, so
don't commit that part:

1. `Create task N in new structure`
   Everything that sets up the new-structure scaffolding in one commit: `README.md` plus any
   extracted data file (`grid.txt`, `number.txt`, etc. from Step 2.3).
2. `Move task N to new structure` (only if there was a legacy file/folder to migrate)
   The ported solution file(s) (e.g. `js/solution_1.js`, updated to read the extracted data file)
   *and* the deletion of the old standalone file/folder — both in this one commit, since it's a
   single conceptual move.
3. *(left to the user)* Their own commit adding the real solution (e.g. `go/solution_N.go` +
   tests) once they've worked through it — don't make this commit yourself.

Only commit when the user asks for it (per the repo-wide git workflow rules).
