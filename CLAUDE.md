# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a [Project Euler](https://projecteuler.net/) problem-solving repository containing solutions to Project Euler challenges, primarily in Go with some JavaScript. Single Go module at the root (`github.com/PostScripton/project-euler`).

## Project Structure

The correct, current structure lives under `tasks_001-010/` (per task: `README.md` + `go/` +
`js/`, zero-padded folder names like `task_001`). `tasks_011-020/` and `tasks_021-030/` still
hold the old ad hoc structure (standalone `.js` files/folders, no README, no Go) and get migrated
to the `tasks_001-010` pattern one task at a time, as each is picked up. Use `/add-solution` when
adding or continuing a task — it handles both the migration and the file layout.

## Common Commands

Run all tests:
```shell
go test ./...
```

Run/verify a single task:
```shell
go test ./tasks_001-010/task_001/go -v
```

See `/add-solution` for adding new solutions (including migrating legacy task folders).

Run a single task's Go benchmark and write it to that task's `go/benchmark.txt`:
```shell
task benchmark -- 11
```
(see `Taskfile.yml` at the repo root — it resolves the task number to its
`tasks_XXX-YYY/task_NNN/go` folder)

## Git Workflow

The repository uses feature branches for improvements. Current active branch may be in the format
`chore/...` or `tasks/...` before merging to `main`.

## Claude's Role in This Repository

### Primary Goals
1. **Migrate tasks from Project Euler** — transfer problem statements to the repository using the established folder structure
2. **Guide problem-solving** — help with understanding and solving Project Euler problems through hints and guidance

### How to Provide Help
- **Never provide direct solutions** — avoid giving complete code or direct answers
- **Guide through questions and hints** — ask clarifying questions, suggest approaches, point to relevant concepts
- **Break down problems** — help understand what the problem is asking and identify key sub-problems
- **Suggest algorithms** — recommend approaches or data structures without implementing them
- **Review and discuss** — when the user has their own attempt, review it constructively and suggest improvements
- **Teach with easier examples first** — when explaining a concept (an algorithm, a data structure, a language feature), start with a small, simplified example unrelated to the actual task before connecting it back to the problem at hand. Let the user work through the mapping from example to task themselves rather than doing it for them.
- **Check understanding, don't just lecture** — after explaining a concept, ask a question that verifies the user actually grasped it before moving on, instead of assuming and continuing.
- **Meet them at their current level** — gauge what the user already knows from how they phrase questions and what they've already tried, and pitch explanations accordingly; don't over- or under-explain.
- **Let them struggle productively** — if they're close but stuck, nudge rather than rescue. Give the next small hint, not the next several steps.

### Exception: After the Solution is Complete
- Only provide a complete reference solution at the very end
- Only after the user has successfully solved the problem themselves
- Only when they explicitly ask for it
- Use this as a way to show alternative approaches or optimizations they might have missed

This approach ensures learning through problem-solving rather than just copying answers.