### Syntax
```go
defer <call>()   // scheduled to run just before the function returns
```

`defer` schedules a function call to run **right before the enclosing function returns** — no matter how it returns (normal return, or even a panic).

```go
func process() {
    fmt.Println("start")
    defer fmt.Println("cleanup")
    fmt.Println("middle")
}
// Output:
// start
// middle
// cleanup
```

`defer` is most often used for cleanup — closing a file, releasing a lock, closing a database connection — right next to the code that opened it, so it's impossible to forget.

> **Gotcha:** Multiple `defer` calls run in **LIFO** order (last deferred, first executed) — like a stack.
> ```go
> defer fmt.Println("A")
> defer fmt.Println("B")
> // prints B then A
> ```

```go
package main

import "fmt"

func main() {
    defer fmt.Println("A")
    defer fmt.Println("B")
    fmt.Println("C")
}
// Output: C, B, A
```