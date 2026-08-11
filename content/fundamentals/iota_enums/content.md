Go doesn't have a built-in `enum` keyword like some languages. Instead, it has `iota` — a special counter that starts at `0` and increments by one for every line inside a `const` block.

```go
const (
    Sunday = iota // 0
    Monday        // 1
    Tuesday       // 2
    Wednesday     // 3
)
```

Each constant automatically gets the next number, so you only have to write `iota` once, on the first line — every following line repeats the same expression with `iota` incremented.

This is Go's idiomatic way to build a set of related, numbered constants — think priority levels, days of the week, or status codes — without manually typing out each number.

> **Gotcha:** `iota` resets to `0` at the start of *every* new `const` block. It only counts within a single block.

```go
package main

import "fmt"

const (
    Low = iota // 0
    Medium     // 1
    High       // 2
)

func main() {
    fmt.Println(Low, Medium, High) // 0 1 2
}
```