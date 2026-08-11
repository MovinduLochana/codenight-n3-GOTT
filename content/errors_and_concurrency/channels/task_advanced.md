> Complete `EvenSum(n int) int`. Launch a goroutine that sends every even number from `1` to `n` (inclusive) into a channel, then **closes** it. Back in `main`, sum the values using `for v := range ch`. Return the total.
>
> **Expected behavior:**
> ```go
> EvenSum(10) // 2+4+6+8+10 = 30
> EvenSum(5)  // 2+4 = 6
> EvenSum(0)  // 0
> ```