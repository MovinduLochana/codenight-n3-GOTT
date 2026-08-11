This lesson introduces **no new concept** — it's a consolidation drill that combines the three Chapter 6 ideas: errors as values (6.1), goroutines + `WaitGroup` (6.2), and channels (6.3) inside one tiny "checkout" app.

```go
func Validate(item string, qty int) error {
    if qty < 1 {
        return errors.New("quantity must be at least 1")
    }
    return nil
}
```

The shape of real Go code you'll start writing next: a function that returns `error`, a worker that runs concurrently, a channel that carries results, and a `WaitGroup` that makes sure every worker is accounted for before the program moves on.

> **Reminder:** channels are for coordination — send on one side, receive on the other, `close(ch)` to signal no more values. Never `close` from the receiving side.

```go
package main

import "fmt"

func Validate(item string, qty int) error {
    if qty < 1 {
        return fmt.Errorf("invalid quantity %d for %s", qty, item)
    }
    return nil
}

func main() {
    fmt.Println(Validate("Coffee", 2)) // <nil>
    fmt.Println(Validate("Coffee", 0)) // invalid quantity 0 for Coffee
}
```