> Complete `MakeCounter() func() int`. Each time the returned function is called, it should return an incrementing count, starting at `1`.
>
> **Expected behavior:**
> ```go
> counter := MakeCounter()
> counter() // 1
> counter() // 2
> counter() // 3
> ```