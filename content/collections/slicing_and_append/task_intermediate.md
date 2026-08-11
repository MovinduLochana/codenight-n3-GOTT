> Complete `RemoveAt(s []int, i int) []int`. It should return a new slice with the element at index `i` removed, keeping the remaining order intact. The idiomatic one-liner is `append(s[:i], s[i+1:]...)`.
>
> > **Gotcha:** this mutates `s`'s backing array — the caller's original slice is changed too. That's expected here; we're demonstrating shared-array behavior, not defending against it.
>
> **Expected behavior:**
> ```go
> RemoveAt([]int{10, 20, 30}, 1) // [10 30]
> RemoveAt([]int{5, 6, 7, 8}, 0) // [6 7 8]
> ```