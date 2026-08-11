Go's `fmt` package gives you three main ways to produce output, and a set of **verbs** to control exactly how values are formatted.

```go
fmt.Println("Age:", 30)               // adds spaces & a newline automatically
fmt.Printf("Age: %d\n", 30)           // format string with explicit verbs
message := fmt.Sprintf("Age: %d", 30) // same as Printf, but returns a string
```

The most common verbs:

| Verb | Meaning |
|------|---------|
| `%v` | default format for any value |
| `%d` | integer |
| `%s` | string |
| `%.2f` | float, rounded to 2 decimal places |
| `%T` | the Go type of the value |

```go
age := 30
fmt.Printf("Age: %d, Type: %T\n", age, age) // Age: 30, Type: int
```

> **Key point:** `Println` is quick and easy for debugging, but `Printf`/`Sprintf` give you precise control over formatting — essential once you start building real output like logs, receipts, or error messages.

```go
package main

import "fmt"

func main() {
    age := 30
    fmt.Printf("Age: %d, Type: %T\n", age, age)
}
```