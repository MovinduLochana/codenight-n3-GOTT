# Chapter 6 — Errors & a Taste of Concurrency
## Content Slides (Concept Screens)

Each slide below is the full-screen markdown shown to the student *before* the corresponding hands-on task screen. Each lesson follows the same structure: **Syntax → Explanation → Example → Callout → Full Runnable Snippet.**

---

## Lesson 6.1 — Errors as Values

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

---

## Lesson 6.2 — Goroutines & WaitGroup

### Syntax
```go
go someFunc()      // run someFunc concurrently, don't block main
var wg sync.WaitGroup
wg.Add(1)          // register one goroutine we'll wait for
go func() { defer wg.Done(); /* work */ }()
wg.Wait()          // block until every Add'd goroutine calls Done
```

A **goroutine** is a lightweight, independently-running function — Go's version of a thread, but far cheaper (you can run thousands of them). Start one by putting `go` before a function call.

```go
go sayHello() // runs concurrently, doesn't block main()
```

Because goroutines run concurrently, you need a way to wait for them to finish before your program moves on (or exits). `sync.WaitGroup` is the standard tool for that:

```go
var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println("worker", id)
    }(i)
}
wg.Wait() // blocks until all 3 goroutines call Done()
```

> **Gotcha:** If multiple goroutines write to the same shared variable at the same time, you get a **data race**. Guard shared state with a `sync.Mutex`, as shown in the tasks below.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("hello from a goroutine")
    }()
    wg.Wait()
}
```

---

## Lesson 6.3 — Channels & Shared State

### Syntax
```go
ch := make(chan int)       // unbuffered channel
ch := make(chan int, 3)    // buffered channel
ch <- v                    // send
v := <-ch                  // receive (blocks until a value arrives)
close(ch)                  // signal "no more values"
for v := range ch { ... }  // drain until channel is closed
```

A **channel** is a typed pipe goroutines use to send values to each other safely — no manual locking required.

```go
ch := make(chan int)

go func() {
    ch <- 42 // send
}()

value := <-ch // receive (blocks until a value arrives)
fmt.Println(value) // 42
```

Closing a channel with `close(ch)` signals "no more values are coming." A `for range` loop over a channel automatically stops once the channel is closed and drained:

```go
go func() {
    for i := 1; i <= 3; i++ {
        ch <- i
    }
    close(ch)
}()

for v := range ch {
    fmt.Println(v) // 1, 2, 3
}
```

> **Gotcha:** Sending on a channel with no receiver (and no buffer space) **blocks forever** — a common cause of deadlocks for Go beginners. Always make sure something is reading.

```go
package main

import "fmt"

func main() {
    ch := make(chan int)
    go func() {
        ch <- 1
        close(ch)
    }()
    for v := range ch {
        fmt.Println(v)
    }
}
```

---

## Lesson 6.4 — Capstone: Concurrent Checkout

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