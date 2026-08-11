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