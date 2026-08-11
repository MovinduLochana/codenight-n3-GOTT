### Syntax
```go
type Outer struct {
    Inner  // embedded without a field name — fields/methods promoted
    Extra  string
}
```

Go doesn't have class inheritance — instead, it uses **composition** through struct embedding. An embedded struct's fields and methods are "promoted" to the outer struct automatically.

```go
type Rectangle struct {
    Width, Height float64
}
func (r Rectangle) Area() float64 { return r.Width * r.Height }

type NamedRectangle struct {
    Rectangle // embedded, no field name
    Name      string
}
```

Because `Rectangle` is embedded, `NamedRectangle` gets its `Area()` method and its `Width`/`Height` fields for free:

```go
nr := NamedRectangle{
    Rectangle: Rectangle{Width: 3, Height: 4},
    Name:      "MyRect",
}
fmt.Println(nr.Area())  // 12 — promoted from Rectangle
fmt.Println(nr.Width)   // 3  — promoted from Rectangle
```

> **Key point:** Embedding is "has-a via promotion," not "is-a" — `NamedRectangle` isn't a `Rectangle` in the type-system sense, but it behaves like one for field/method access. This is Go's alternative to inheritance: prefer composition.

```go
package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}
func (r Rectangle) Area() float64 { return r.Width * r.Height }

type NamedRectangle struct {
    Rectangle
    Name string
}

func main() {
    nr := NamedRectangle{Rectangle: Rectangle{Width: 3, Height: 4}, Name: "MyRect"}
    fmt.Println(nr.Name, nr.Area()) // MyRect 12
}
```