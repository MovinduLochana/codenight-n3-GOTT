> Complete the function `Report(name string, orders int, revenue float64) string`. It should build a **multi-line** string with `fmt.Sprintf` using `\n` escapes in this exact format:
>
> ```
> Customer: <name>
> Orders: <orders>
> Revenue LKR <revenue to 2 decimal places>
> ```
>
> **Expected behavior:**
> ```go
> Report("Nimal", 12, 3450.5)
> // "Customer: Nimal\nOrders: 12\nRevenue LKR 3450.50"
> ```