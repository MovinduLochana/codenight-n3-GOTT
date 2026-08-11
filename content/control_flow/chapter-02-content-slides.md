# Chapter 2 — Control Flow
## Content Slides (Concept Screens)

Each slide below is the full-screen markdown shown to the student *before* the corresponding hands-on task screen. Each lesson follows the same structure: **Syntax → Explanation → Example → Callout → Full Runnable Snippet.**

---

## Lesson 2.1 — `if`, `else`, and Init Statements

### Syntax
```go
if <condition> {
    // ...
} else if <condition> {
    // ...
} else {
    // ...
}

if <init>; <condition> {
    // ...
}
```

Go's `if` statements don't need parentheses around the condition, but the curly braces are mandatory — even for a single line. You can chain an `else if` for more branches and finish with an `else` catch-all.

```go
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("F")
}
```

Go also lets you run a short statement *before* the condition, scoped only to the `if`/`else` block — extremely common for error-checking:

```go
if err := doSomething(); err != nil {
    fmt.Println("failed:", err)
}
```

> **Key point:** The variable declared in an `if`-init statement (like `err` above) only exists inside that `if`/`else` chain — it disappears afterward. This keeps temporary variables from leaking into the rest of the function.

```go
package main

import "fmt"

func main() {
    score := 85
    if score >= 90 {
        fmt.Println("A")
    } else if score >= 80 {
        fmt.Println("B")
    } else {
        fmt.Println("F")
    }
}
```

---

## Lesson 2.2 — The One True Loop: `for`

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

---

## Lesson 2.3 — `switch`

### Syntax
```go
switch <expression> {
case <value1>, <value2>:
    // ...
default:
    // ...
}

switch { // expression-less form
case <condition>:
    // ...
default:
    // ...
}
```

A `switch` in Go compares a value against several cases — but unlike C or Java, **it does not fall through by default**. Only the matching case runs.

```go
switch day {
case "Sat", "Sun":
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}
```

You can also omit the expression entirely, turning `switch` into a clean replacement for a long `if`/`else if` chain:

```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
default:
    fmt.Println("F")
}
```

> **Gotcha:** Coming from C or Java, you might expect to need `break` after each case — in Go it's automatic. Use the `fallthrough` keyword if you explicitly want the old behavior (rare).

```go
package main

import "fmt"

func main() {
    day := "Sat"
    switch day {
    case "Sat", "Sun":
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
}
```
