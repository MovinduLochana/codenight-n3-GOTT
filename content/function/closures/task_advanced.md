> Complete `MakeBank(balance float64) func(amount float64) float64`. The returned closure captures `balance` and every call **deposits** `amount` (if positive) and returns the new running balance. If `amount` is negative or zero, leave the balance unchanged and return the current balance. Two independent banks must never share state.
>
> **Expected behavior:**
> ```go
> bank := MakeBank(100)
> bank(50)  // 150
> bank(-20) // 150 (ignored)
> bank(25)  // 175
> ```