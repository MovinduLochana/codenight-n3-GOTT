> Complete `ConcurrentSum(nums []int) int`. Launch a goroutine for **each** number; each goroutine adds its number to a shared `total` guarded by a mutex. A `sync.WaitGroup` waits for all of them. Return the final total. Order does not matter — only correctness under concurrency.
>
> **Expected behavior:**
> ```go
> ConcurrentSum([]int{1, 2, 3, 4}) // 10
> ConcurrentSum([]int{})           // 0
> ```