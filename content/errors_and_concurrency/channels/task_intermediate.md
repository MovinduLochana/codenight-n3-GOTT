> Complete `SquaresPipeline(k int) []int`. Build a two-goroutine **pipeline**: a producer sends `1..k` into a first channel, a transform goroutine reads it, squares each value, and sends the result into a second channel (then closes it). `main` drains the second channel with `for v := range ch` into the returned slice.
>
> **Expected behavior:**
> ```go
> SquaresPipeline(4) // [1 4 9 16]
> SquaresPipeline(0) // []
> ```