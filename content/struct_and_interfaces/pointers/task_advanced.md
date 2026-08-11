> Complete `SwapNums(a, b *int)` so it exchanges the values pointed to by `a` and `b` — no return value needed. This is the classic "how would you do this without pointers?" problem.
>
> **Expected behavior:**
> ```go
> x, y := 1, 2
> SwapNums(&x, &y)
> x // 2
> y // 1
> ```