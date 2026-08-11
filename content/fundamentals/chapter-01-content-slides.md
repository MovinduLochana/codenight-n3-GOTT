# Chapter 1 — Go Fundamentals
## Content Slides (Concept Screens)

Each slide below is the full-screen markdown shown to the student *before* the corresponding hands-on task screen. Each lesson now follows the same structure: **Syntax → Explanation → Example → Callout → Full Runnable Snippet.**

---

## Lesson 1.1 — Declaring Variables with `var`

### Syntax
```go
var <name> <type> = <value>
var <name> = <value>          // type inferred
var <name1>, <name2> <type> = <value1>, <value2>
```

In Go, you declare a variable using the `var` keyword, followed by its name, its type, and (optionally) a starting value.

```go
var name string = "Go"
var year int = 2009
```

If you provide a value, Go can usually **infer** the type for you, so you can drop it:

```go
var name = "Go"   // inferred as string
var year = 2009   // inferred as int
```

You can also declare several variables of the same type in one line:

```go
var a, b int = 1, 2
```

> **Key point:** `var` is the explicit, "spell it all out" way to declare a variable. It works both inside functions and at the package level (outside any function) — which matters because Go's other declaration style, `:=`, does not.

```go
package main

import "fmt"

func main() {
    var language string = "Go"
    var releaseYear = 2009 // type inferred as int

    fmt.Println(language, releaseYear)
}
```

---

## Lesson 1.2 — Short Declaration `:=`

### Syntax
```go
<name> := <value>
<name1>, <name2> := <value1>, <value2>
```

Inside a function, Go gives you a shortcut for declaring and initializing a variable in one step: the `:=` operator. It infers the type automatically, so you never write the word `var` or the type name.

```go
name := "Go"
year := 2009
```

This is completely equivalent to:

```go
var name = "Go"
var year = 2009
```

`:=` is by far the most common way you'll declare variables inside Go functions — it's shorter and just as type-safe, since Go still knows `name` is a `string` and `year` is an `int` behind the scenes.

> **Gotcha:** `:=` only works **inside functions**, not at the package level. It also requires at least one *new* variable on the left-hand side — you can't use it to just reassign existing variables.

```go
package main

import "fmt"

func main() {
    language := "Go"
    releaseYear := 2009

    fmt.Println(language, releaseYear)
}
```

---

## Lesson 1.3 — Zero Values

### Syntax
```go
var <name> <type>   // declared with no value — gets the zero value automatically
```

What happens if you declare a variable but don't give it a value? Unlike some languages, Go doesn't leave it as garbage or `undefined`. Every type has a **zero value** — a sensible default it's automatically set to.

| Type | Zero value |
|------|-----------|
| `int`, `float64` | `0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| pointers, slices, maps, etc. | `nil` |

```go
var score int
var passed bool

fmt.Println(score, passed) // 0 false
```

> **Key point:** Zero values mean a Go variable is *always* in a valid, usable state the moment it's declared — you never have to worry about reading "uninitialized memory."

```go
package main

import "fmt"

func main() {
    var count int
    var message string
    var isReady bool

    fmt.Println(count, message, isReady) // 0 "" false
}
```

---

## Lesson 1.4 — Basic Types Tour

### Syntax
```go
var <name> int       = <whole number>
var <name> float64   = <decimal number>
var <name> string    = "<text>"
var <name> bool      = true | false
```

Go has a rich set of built-in types, but four cover almost everything you'll need to get started:

| Type | Used for | Example literal |
|------|----------|-----------------|
| `int` | whole numbers | `42` |
| `float64` | decimal numbers | `3.14` |
| `string` | text | `"hello"` |
| `bool` | true/false logic | `true` |

```go
var age int = 30
var price float64 = 19.99
var city string = "Colombo"
var isOpen bool = true
```

Go is **statically typed**: once a variable is declared as one of these types, it can never hold a value of a different type.

```go
var age int = 30
age = "thirty" // compile error: cannot use "thirty" (string) as int
```

> **Note:** Go also has more specific numeric types (`int8`, `int64`, `uint`, `float32`, etc.) for when you need precise control over memory or range — we won't need them in this crash course, but they're worth exploring later.

```go
package main

import "fmt"

func main() {
    var age int = 30
    var price float64 = 19.99
    var city string = "Colombo"
    var isOpen bool = true

    fmt.Println(age, price, city, isOpen)
}
```

---

## Lesson 1.5 — Constants (`const`)

### Syntax
```go
const <name> = <value>

const (
    <name1> = <value1>
    <name2> = <value2>
)
```

A **constant** is a value that's fixed at compile time and can never be reassigned. You declare one with the `const` keyword.

```go
const Pi = 3.14
const MaxUsers = 100
```

Trying to change a constant is a compile-time error, not something you discover at runtime:

```go
const MaxUsers = 100
MaxUsers = 200 // compile error: cannot assign to MaxUsers
```

You can group related constants together using a `const` block:

```go
const (
    StatusActive   = "active"
    StatusInactive = "inactive"
)
```

> **Key point:** Use `const` for values that should never change — configuration limits, fixed labels, mathematical constants. It documents intent and lets the compiler catch accidental reassignment for you.

```go
package main

import "fmt"

func main() {
    const MaxScore = 100
    fmt.Println(MaxScore)
}
```

---

## Lesson 1.6 — `iota` for Enums

### Syntax
```go
const (
    <name1> = iota   // 0
    <name2>          // 1
    <name3>          // 2
    // ... increments by 1 each line
)
```

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

---

## Lesson 1.7 — Operators

### Syntax
```go
<a> + <b>   <a> - <b>   <a> * <b>   <a> / <b>   <a> % <b>   // arithmetic
<a> == <b>  <a> != <b>  <a> < <b>   <a> > <b>   <a> <= <b>  <a> >= <b>  // comparison
<a> && <b>  <a> || <b>  !<a>                                // logical
```

Go supports the same three families of operators you'd expect from most languages:

**Arithmetic:** `+` `-` `*` `/` `%` (modulo — remainder after division)

```go
sum := 5 + 3       // 8
remainder := 7 % 2 // 1
```

**Comparison:** `==` `!=` `<` `>` `<=` `>=` — always return a `bool`

```go
isEqual := (5 == 5) // true
```

**Logical:** `&&` (and), `||` (or), `!` (not)

```go
canVote := age >= 18 && hasID
```

> **Gotcha:** Integer division **truncates** — it drops the decimal part entirely, it doesn't round.
> ```go
> result := 7 / 2 // 3, not 3.5
> ```
> We'll fix this properly in the next lesson on type conversion.

```go
package main

import "fmt"

func main() {
    n := 7
    isEven := n%2 == 0
    fmt.Println(isEven) // false
}
```

---

## Lesson 1.8 — Type Conversion

### Syntax
```go
<targetType>(<value>)
// e.g. float64(x), int(y), string(z)
```

Go **never** converts types for you automatically — not even between closely related numeric types like `int` and `float64`. Every conversion must be explicit.

```go
var i int = 7
var f float64 = float64(i) // explicit conversion
```

This matters most with division. Two `int`s divided by each other produce a truncated `int` result:

```go
a, b := 7, 2
result := a / b // 3 — the .5 is silently dropped!
```

To get a true decimal result, convert *before* dividing:

```go
result := float64(a) / float64(b) // 3.5
```

> **Key point:** The conversion syntax looks like a function call — `float64(x)`, `int(y)`, `string(z)` — but it's a compiler-level cast, not a runtime function. It works between compatible types only (you can't `int("hello")`, for instance).

```go
package main

import "fmt"

func main() {
    a, b := 7, 2
    result := float64(a) / float64(b)
    fmt.Println(result) // 3.5
}
```

---

## Lesson 1.9 — `fmt` Print Verbs

### Syntax
```go
fmt.Println(<values...>)
fmt.Printf("<format string with verbs>\n", <values...>)
<result> := fmt.Sprintf("<format string with verbs>", <values...>)
```

Go's `fmt` package gives you three main ways to produce output, and a set of **verbs** to control exactly how values are formatted.

```go
fmt.Println("Age:", 30)               // adds spaces & a newline automatically
fmt.Printf("Age: %d\n", 30)           // format string with explicit verbs
message := fmt.Sprintf("Age: %d", 30) // same as Printf, but returns a string
```

The most common verbs:

| Verb | Meaning |
|------|---------|
| `%v` | default format for any value |
| `%d` | integer |
| `%s` | string |
| `%.2f` | float, rounded to 2 decimal places |
| `%T` | the Go type of the value |

```go
age := 30
fmt.Printf("Age: %d, Type: %T\n", age, age) // Age: 30, Type: int
```

> **Key point:** `Println` is quick and easy for debugging, but `Printf`/`Sprintf` give you precise control over formatting — essential once you start building real output like logs, receipts, or error messages.

```go
package main

import "fmt"

func main() {
    age := 30
    fmt.Printf("Age: %d, Type: %T\n", age, age)
}
```

---

## Lesson 1.10 — String Formatting Practice

### Syntax
```go
<result> := fmt.Sprintf("<literal text> %s <literal text> %d <literal text> %.2f", <stringVal>, <intVal>, <floatVal>)
```

You've now seen `var`, `:=`, zero values, the basic types, constants, `iota`, operators, type conversion, and `fmt`'s print verbs. This lesson doesn't introduce anything new — it's about combining what you've learned to build a real, formatted string using `Sprintf`, something you'll do constantly in Go: building log lines, error messages, and CLI output.

```go
item := "Coffee"
qty := 2
price := 4.50

receipt := fmt.Sprintf("Item: %s, Qty: %d, Price: $%.2f", item, qty, price)
fmt.Println(receipt) // Item: Coffee, Qty: 2, Price: $4.50
```

Notice how each verb lines up with the type of the value passed in: `%s` for the string, `%d` for the int, `%.2f` for the float.

> **Reminder:** `Sprintf` doesn't print anything — it *returns* a string. You still need `Println` or `Print` if you want it to show up on screen.

```go
package main

import "fmt"

func main() {
    item := "Coffee"
    qty := 2
    price := 4.50

    receipt := fmt.Sprintf("Item: %s, Qty: %d, Price: $%.2f", item, qty, price)
    fmt.Println(receipt)
}
```
