### Syntax
```go
func <name>(<params>) (<ret1>, <ret2>) { ... }
func <name>(<params>) (<name1, name2 T>) { ... } // named returns
```

Go functions can return more than one value — most commonly used for the `(result, error)` pattern you'll see everywhere in Go code.

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

Callers typically check the error immediately:

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("error:", err)
}
```

> **Key point:** Returning `(value, error)` instead of throwing exceptions is Go's core idiom for error handling — you'll use this pattern constantly, so get comfortable with it now. We'll cover `error` in depth in Chapter 6.

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    fmt.Println(result, err) // 5 <nil>
}
```