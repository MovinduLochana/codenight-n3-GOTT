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