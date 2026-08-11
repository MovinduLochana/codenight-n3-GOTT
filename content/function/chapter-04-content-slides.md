# Chapter 4 — Functions
## Content Slides (Concept Screens)

Each slide below is the full-screen markdown shown to the student *before* the corresponding hands-on task screen. Each lesson follows the same structure: **Syntax → Explanation → Example → Callout → Full Runnable Snippet.**

---

## Lesson 4.1 — Functions & Multiple Return Values

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

---

## Lesson 4.2 — Variadic Functions

### Syntax
```go
func <name>(<params>, last ...T) { ... } // last param is a []T
sum(slice...)                            // spread a slice into a variadic call
```

A **variadic** function accepts a variable number of arguments of the same type — you've already used one every time you called `fmt.Println`.

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)       // works with any number of args
sum()              // even zero args
```

Inside the function, `nums` behaves exactly like a regular `[]int` slice. You can also "spread" an existing slice into a variadic call using `...`:

```go
values := []int{1, 2, 3}
sum(values...)
```

> **Key point:** A variadic parameter must be the **last** parameter in the function signature, and a function can only have one.

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3)) // 6
}
```

---

## Lesson 4.3 — Closures

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

---

## Lesson 4.4 — `defer`

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