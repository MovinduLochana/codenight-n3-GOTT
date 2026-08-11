### Syntax
```go
for init; condition; post { ... } // classic C-style loop
for condition { ... }             // "while" loop
for { ... }                       // infinite loop
```

Go has exactly **one** looping keyword: `for`. It covers every case other languages split across `for`, `while`, and `do-while`.

```go
// classic C-style loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// "while" loop — condition only
n := 0
for n < 3 {
    n++
}

// infinite loop — needs an explicit break
for {
    break
}
```

Use `continue` to skip to the next iteration, and `break` to exit the loop entirely.

> **Key point:** There's no `while` keyword in Go — reach for `for` with just a condition whenever you'd normally write `while` in another language.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        if i == 1 {
            continue
        }
        fmt.Println(i)
    }
}
```