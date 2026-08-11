> Complete `Squares(nums []int) []int`. Every number should be squared **in its own goroutine** and the result stored into a shared result slice guarded by a mutex. Because goroutine completion order is unpredictable, **sort** the result before returning so the output is deterministic. Return a slice of the same length as `nums`, where every element is `nums[i] * nums[i]`.
>
> **Expected behavior (order guaranteed after sorting):**
> ```go
> Squares([]int{4, 1, 3}) // [1 9 16]
> Squares([]int{})        // []
> ```