> Complete `Prices(totals chan int)`. It should launch a goroutine that sends the totals `10`, `20`, `30` (in order) into the `totals` channel and then **closes** it — the classic producer. Wait for it with a `sync.WaitGroup`. Back in `main`, sum the values with `for v := range totals` and return the total.
>
> **Expected behavior:**
> ```go
> CheckoutTotals() // 60
> ```