> Complete the function `Receipt(item string, qty int, price float64) string`. It should return a string built with `fmt.Sprintf` in this exact format:
> ```
> Item: <item>, Qty: <qty>, Price: $<price to 2 decimal places>
> ```
>
> **Expected behavior:**
> ```go
> Receipt("Coffee", 2, 4.50) // "Item: Coffee, Qty: 2, Price: $4.50"
> ```