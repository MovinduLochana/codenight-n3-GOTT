> Implement a function `MakeFilter(limit int) func([]int) []int` that returns a closure.
> The closure should accept a slice of integers and return a new slice containing only the elements that are **strictly greater** than the `limit` configured when the closure was created.
>
> **Expected Output:**
>
> ```
> Above 10: [20 15 30]
> Above 20: [30]
> ```
