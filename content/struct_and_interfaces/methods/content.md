### Syntax
```go
func (r TypeName) Method() T { ... }     // value receiver — works on a copy
func (r *TypeName) Method() T { ... }    // pointer receiver — can mutate original
```

A **method** is a function attached to a specific type via a **receiver**. There are two kinds of receivers:

```go
// value receiver — gets a copy, can't mutate the original
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// pointer receiver — gets a reference, CAN mutate the original
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

```go
r := Rectangle{Width: 3, Height: 4}
r.Scale(2) // Go automatically takes &r here
fmt.Println(r.Width) // 6
```

> **Gotcha:** If a method needs to modify the receiver's fields (or the struct is large and copying it is wasteful), use a **pointer receiver** (`*Type`). If it only reads data, a value receiver is fine.

```go
package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    r := Rectangle{Width: 3, Height: 4}
    r.Scale(2)
    fmt.Println(r.Width, r.Height) // 6 8
}
```