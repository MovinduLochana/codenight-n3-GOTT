> Complete `InsertAt(s []int, i, v int) []int`. It should return a new slice identical to `s` but with `v` inserted at index `i`, preserving order. Idiomatic approach: lengthen `s` by one with `append`, shift the tail right with `copy`, then set index `i` to `v`. Assume `i` is a valid index and `s` is non-empty.
>
> **Expected behavior:**
> ```go
> InsertAt([]int{10, 20, 30}, 1, 99) // [10 99 20 30]
> InsertAt([]int{1, 2}, 0, 0)        // [0 1 2]
> ```