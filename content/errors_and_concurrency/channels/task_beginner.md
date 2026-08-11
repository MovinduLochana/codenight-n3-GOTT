> Complete `Producer(ch chan<- int, n int)`. It should send the numbers `1` through `n` (inclusive) into `ch`, in order, then close the channel. The consumer code in `main` is already written for you.
>
> **Expected behavior:**
> ```go
> // Producer(ch, 3) sends 1, 2, 3 then closes ch
> // consumer prints: 1 2 3
> ```