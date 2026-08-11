### Syntax
```go
if <condition> {
    // ...
} else if <condition> {
    // ...
} else {
    // ...
}

if <init>; <condition> {
    // ...
}
```

Go's `if` statements don't need parentheses around the condition, but the curly braces are mandatory — even for a single line.

```go
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("F")
}
```

Go also lets you run a short statement *before* the condition, scoped only to the `if`/`else` block — extremely common for error-checking:

```go
if err := doSomething(); err != nil {
    fmt.Println("failed:", err)
}
```

> **Key point:** The variable declared in an `if`-init statement (like `err` above) only exists inside that `if`/`else` chain — it disappears afterward. This keeps temporary variables from leaking into the rest of the function.

```go
package main

import "fmt"

func main() {
    score := 85
    if score >= 90 {
        fmt.Println("A")
    } else if score >= 80 {
        fmt.Println("B")
    } else {
        fmt.Println("F")
    }
}
```