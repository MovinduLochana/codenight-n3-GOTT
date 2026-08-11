> Complete `IndexAll(nums []int, target int) []int`. It should return a slice holding the index of every position where `target` appears. Build the result with `append` inside a `range` loop that uses the index (`i`).
>
> **Expected behavior:**
> ```go
> IndexAll([]int{1, 2, 1, 4, 1}, 1) // [0 2 4]
> IndexAll([]int{1, 2, 3}, 9)       // []
> ```