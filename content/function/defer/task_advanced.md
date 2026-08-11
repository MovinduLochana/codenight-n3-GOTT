> Complete `LoopCleanup(n int) (nums []int)`. Inside a `for i := 0; i < n; i++` loop, defer an append of `i` to the named return. Because deferred calls run **after** the loop, in LIFO order, the collected numbers come back **reversed**: `n-1` down to `0`.
>
> **Expected behavior (Go 1.22+ per-iteration loop semantics):**
> ```go
> LoopCleanup(3) // [2 1 0]
> LoopCleanup(1) // [0]
> ```