### Syntax
```go
s[low:high]  // sub-slice from index low up to (not including) high
copy(dst, src) // copies elements into an independent slice
```

You can carve out a sub-slice using `slice[low:high]`:

```go
nums := []int{10, 20, 30, 40, 50}
middle := nums[1:3] // [20 30]
```

> **Gotcha:** A sub-slice shares the **same underlying array** as the original. Modifying one can silently modify the other — a classic source of bugs.
> ```go
> a := []int{1, 2, 3}
> b := a[0:2]
> b[0] = 99
> fmt.Println(a) // [99 2 3] — a changed too!
> ```
> When you need an independent copy, use `copy()`:
> ```go
> b := make([]int, len(a))
> copy(b, a)
> ```

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30, 40, 50}
    middle := nums[1:3]
    fmt.Println(middle) // [20 30]
}
```