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