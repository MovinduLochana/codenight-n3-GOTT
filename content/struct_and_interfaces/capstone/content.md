This lesson introduces **no new concept** — it's a consolidation drill that combines structs, methods, interfaces, and embedding from Lessons 5.1–5.5 into one tiny café menu app.

```go
type Product struct {
    Name  string
    Price float64
}

func (p Product) Label() string {
    return fmt.Sprintf("%s: $%.2f", p.Name, p.Price)
}
```

The pattern you'll keep using: a struct holds the data, methods attach behavior, an interface lets different kinds of products be treated uniformly, and embedding lets you extend a product without re-declaring its fields.

> **Reminder:** `Sprintf` builds a string, it doesn't print. Combine it with the `%s` and `%.2f` verbs you learned in Chapter 1.

```go
package main

import "fmt"

type Product struct {
    Name  string
    Price float64
}

func (p Product) Label() string {
    return fmt.Sprintf("%s: $%.2f", p.Name, p.Price)
}

func main() {
    p := Product{Name: "Coffee", Price: 4.50}
    fmt.Println(p.Label()) // Coffee: $4.50
}
```