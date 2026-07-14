Run Go benchmarks for this repository.

Usage: `/benchmark` (all tasks) or `/benchmark <task number>` (single task).

- All packages, saved to file:
  ```shell
  go test -bench=. -benchmem | cat > benchmark.txt
  ```
- Single task (writes to that task's own `go/benchmark.txt`, matching the existing per-task files):
  ```shell
  go test ./tasks_0XX-0YY/task_0NN/go -bench=. -benchmem | cat > tasks_0XX-0YY/task_0NN/go/benchmark.txt
  ```

Resolve `tasks_0XX-0YY/task_0NN` from the given task number the same way `/add-solution` does.
