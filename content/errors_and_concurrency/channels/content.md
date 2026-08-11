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