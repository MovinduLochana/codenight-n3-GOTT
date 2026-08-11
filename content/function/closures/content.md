### Syntax
```go
func <outer>(...) func(<args>) <ret> {
    state := ...          // captured variable
    return func(<args>) <ret> {
        // ... uses & mutates state
    }
}
```

A **closure** is a function that "closes over" — captures and remembers — variables from the scope it was created in, even after that outer function has returned.

```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

counter := makeCounter()
fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
fmt.Println(counter()) // 3
```

Each call to `makeCounter()` creates a **brand new** `count` variable, so separate counters don't interfere with each other.

> **Key point:** Closures are how Go achieves patterns other languages need classes for — stateful function generators, middleware, and callback handlers all lean on this.

```go
package main

import "fmt"

func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter := makeCounter()
    fmt.Println(counter())
    fmt.Println(counter())
}
```