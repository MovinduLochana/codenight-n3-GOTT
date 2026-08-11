### Syntax
```go
switch <expression> {
case <value1>, <value2>:
    // ...
default:
    // ...
}

switch { // expression-less form
case <condition>:
    // ...
default:
    // ...
}
```

A `switch` in Go compares a value against several cases — but unlike C or Java, **it does not fall through by default**. Only the matching case runs.

```go
switch day {
case "Sat", "Sun":
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}
```

You can also omit the expression entirely, turning `switch` into a clean replacement for a long `if`/`else if` chain:

```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
default:
    fmt.Println("F")
}
```

> **Gotcha:** Coming from C or Java, you might expect to need `break` after each case — in Go it's automatic. Use the `fallthrough` keyword if you explicitly want the old behavior (rare).

```go
package main

import "fmt"

func main() {
    day := "Sat"
    switch day {
    case "Sat", "Sun":
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
}
```