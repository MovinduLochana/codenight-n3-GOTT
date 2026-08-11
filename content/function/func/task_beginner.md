> Complete `Divide(a, b int) (int, error)`. Return an error (using `errors.New`) if `b` is `0`; otherwise return `a / b` and `nil`.
>
> **Expected behavior:**
> ```go
> Divide(10, 2) // 5, nil
> Divide(5, 0)  // 0, error
> ```