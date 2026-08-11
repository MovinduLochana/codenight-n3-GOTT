> Complete `Double(ptr *int) int`. It should multiply the value pointed to by `ptr` by 2 (writing back through the pointer), and return the new value.
>
> **Expected behavior:**
> ```go
> x := 7
> Double(&x) // 14
> x          // 14 — mutated through the pointer
> ```