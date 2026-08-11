> Complete the function `Row(item string, price float64) string` that builds a single **right-aligned table row** using width flags in the format string. Use `%-12s` to left-align the item in a 12-char field, then `%10.2f` to right-align the float in a 10-char field, with the fixed separator ` | Price: ` between them:
>
> ```go
> fmt.Sprintf("Item: %-12s | Price: %10.2f", item, price)
> ```
>
> **Expected behavior (exact strings):**
> ```go
> Row("Coffee", 450.5)  // "Item: Coffee       | Price:     450.50"
> Row("Short", 25.0)    // "Item: Short        | Price:      25.00"
> ```
>
> Stretch hint: Go keeps the total field width exact, so shorter items get extra trailing spaces to fill 12 chars.