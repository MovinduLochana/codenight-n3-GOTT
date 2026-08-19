> Fill in the `AverageUntilZero` function so it reads integers one at a time in a loop using `fmt.Scan`, and keeps reading until the user enters `0`. It should return the average (as a `float64`) of all the numbers entered **before** the `0`.
>
> **Example:**
> Input: `10 20 30 0` → `(10+20+30)/3 = 20`
>
> **Hints:**
> - Loop with `for { ... }` and `var n int`.
> - `fmt.Scan(&n)` returns how many values it successfully read — stop the loop if it's `0` (no more input).
> - `if n == 0 { break }` when the user ends the sequence.
> - Track a running `count` and `sum`, then return `float64(sum) / float64(count)`.