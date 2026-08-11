### Syntax
```go
type <Name> interface {
    Method1() Ret
    Method2(Arg) Ret
}
```

An **interface** defines a set of methods a type must have — but unlike many languages, Go types satisfy an interface **implicitly**. There's no `implements` keyword; if the methods match, it just works.

```go
type Shape interface {
    Area() float64
}

type Rectangle struct{ Width, Height float64 }
func (r Rectangle) Area() float64 { return r.Width * r.Height }

type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return 3.14159 * c.Radius * c.Radius }
```

Both `Rectangle` and `Circle` automatically satisfy `Shape` — no declaration needed. This lets you write functions that work with *any* shape:

```go
func TotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, s := range shapes {
        total += s.Area()
    }
    return total
}
```

> **Key point:** This is how Go achieves polymorphism without inheritance — small, implicit interfaces defined by behavior, not by a type hierarchy.

```go
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Rectangle struct{ Width, Height float64 }
func (r Rectangle) Area() float64 { return r.Width * r.Height }

func main() {
    var s Shape = Rectangle{Width: 3, Height: 4}
    fmt.Println(s.Area()) // 12
}
```