Complete `Tail(s []int, n int) []int`. It should return a **new independent copy** of the last `n` elements of `s` (without leaking extra capacity into the result). Use `make` + `copy`. If `n >= len(s)`, return a copy of the whole slice.

**Expected behavior:**
```go
Tail([]int{1, 2, 3, 4, 5}, 2) // [4 5]
Tail([]int{1, 2}, 5)          // [1 2]
Tail([]int{1, 2, 3}, 0)       // []
```