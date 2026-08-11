> Complete `Checkout(items []string, qty []int) ([]int, error)`. It should validate every `(item, qty)` pair with `Validate` (from the beginner task): if any quantity is invalid, return `nil, error` immediately. Otherwise compute each `qty[i] * 2` (a fake per-item price) **in its own goroutine**, send the results through a channel, and collect them into a slice with `for v := range ch`. A `WaitGroup` coordinates the workers. **Sort** the collected slice before returning so the output is deterministic despite goroutine ordering.
>
> **Expected behavior:**
> ```go
> Checkout([]string{"Coffee", "Tea"}, []int{2, 3}) // [4 6], nil
> Checkout([]string{"Coffee"}, []int{0})           // nil, error
> ```