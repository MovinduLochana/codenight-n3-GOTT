> Complete `Contains(nums []int, target int) bool`. It should return `true` if `target` appears anywhere in `nums`. Use a `range` loop that discards the index with `_` and returns early on a match.
>
> **Expected behavior:**
> ```go
> Contains([]int{1, 2, 3}, 2) // true
> Contains([]int{1, 2, 3}, 9) // false
> Contains([]int{}, 0)        // false
> ```