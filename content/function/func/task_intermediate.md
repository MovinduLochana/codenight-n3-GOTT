> Complete `DivMod(a, b int) (quotient, remainder int, err error)`. It should return three values: the quotient and remainder of `a / b` (integer division and `%`), plus an error if `b == 0`. Use Go's **named return values** so you can return them individually (`return 0, 0, errors.New(...)`).
>
> **Expected behavior:**
> ```go
> q, r, err := DivMod(10, 3) // 3, 1, nil
> q, r, err := DivMod(10, 0) // 0, 0, error
> ```