### Syntax
```go
errors.New("message")       // a plain sentinel-less error
fmt.Errorf("msg: %w", err)  // wrap an existing error, preserving it
errors.Is(err, target)      // true if err wraps target
_, ok := err.(SomeType)     // type assertion for custom errors
```

Go doesn't use exceptions. Errors are ordinary values that implement the built-in `error` interface, and functions that can fail simply return one alongside their result.

```go
import "errors"

func parsePositive(s string) (int, error) {
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("invalid number: %w", err)
    }
    if n < 0 {
        return 0, errors.New("number must be positive")
    }
    return n, nil
}
```

Callers check the error immediately after the call, before touching the result:

```go
n, err := parsePositive("-5")
if err != nil {
    fmt.Println("error:", err)
    return
}
```

> **Key point:** `%w` in `fmt.Errorf` "wraps" the original error, preserving it so callers can inspect the underlying cause with `errors.Is`/`errors.As` later — useful once your programs grow beyond this crash course.

```go
package main

import (
    "errors"
    "fmt"
)

func checkAge(age int) error {
    if age < 0 {
        return errors.New("age cannot be negative")
    }
    return nil
}

func main() {
    err := checkAge(-1)
    fmt.Println(err) // age cannot be negative
}
```