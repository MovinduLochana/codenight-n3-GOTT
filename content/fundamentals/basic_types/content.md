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