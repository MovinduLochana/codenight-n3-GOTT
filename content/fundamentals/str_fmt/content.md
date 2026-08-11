You've now seen `var`, `:=`, zero values, the basic types, constants, `iota`, operators, type conversion, and `fmt`'s print verbs. This lesson doesn't introduce anything new — it's about combining what you've learned to build a real, formatted string using `Sprintf`, something you'll do constantly in Go: building log lines, error messages, and CLI output.

```go
item := "Coffee"
qty := 2
price := 4.50

receipt := fmt.Sprintf("Item: %s, Qty: %d, Price: $%.2f", item, qty, price)
fmt.Println(receipt) // Item: Coffee, Qty: 2, Price: $4.50
```

Notice how each verb lines up with the type of the value passed in: `%s` for the string, `%d` for the int, `%.2f` for the float.

> **Reminder:** `Sprintf` doesn't print anything — it *returns* a string. You still need `Println` or `Print` if you want it to show up on screen.

```go
package main

import "fmt"

func main() {
    item := "Coffee"
    qty := 2
    price := 4.50

    receipt := fmt.Sprintf("Item: %s, Qty: %d, Price: $%.2f", item, qty, price)
    fmt.Println(receipt)
}
```