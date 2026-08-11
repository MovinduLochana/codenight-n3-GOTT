A **constant** is a value that's fixed at compile time and can never be reassigned. You declare one with the `const` keyword.

```go
const Pi = 3.14
const MaxUsers = 100
```

Trying to change a constant is a compile-time error, not something you discover at runtime:

```go
const MaxUsers = 100
MaxUsers = 200 // compile error: cannot assign to MaxUsers
```

You can group related constants together using a `const` block:

```go
const (
    StatusActive   = "active"
    StatusInactive = "inactive"
)
```

> **Key point:** Use `const` for values that should never change — configuration limits, fixed labels, mathematical constants. It documents intent and lets the compiler catch accidental reassignment for you.

```go
package main

import "fmt"

func main() {
    const MaxScore = 100
    fmt.Println(MaxScore)
}
```