# Chapter 5 — Structs & Interfaces
## Content Slides (Concept Screens)

Each slide below is the full-screen markdown shown to the student *before* the corresponding hands-on task screen. Each lesson follows the same structure: **Syntax → Explanation → Example → Callout → Full Runnable Snippet.**

---

## Lesson 5.1 — Structs

### Syntax
```go
type <Name> struct {
    Field1 Type1
    Field2 Type2
}

v := <Name>{Field1: val1, Field2: val2} // by name
w := <Name>{val1, val2}                 // positionally (field order)
v.Field1                                // access a field
```

A **struct** groups related fields into a single custom type — Go's building block for modeling data (there are no classes in Go).

```go
type Rectangle struct {
    Width  float64
    Height float64
}

r := Rectangle{Width: 3, Height: 4}
fmt.Println(r.Width) // 3
```

You can also create struct values positionally (matching field order), though named fields are usually clearer:

```go
r := Rectangle{3, 4} // Width: 3, Height: 4
```

> **Key point:** Structs are plain data containers. Behavior (functions that operate on them) is added separately via **methods** — covered in Lesson 5.3.

```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

func Area(r Rectangle) float64 {
    return r.Width * r.Height
}

func main() {
    r := Rectangle{Width: 3, Height: 4}
    fmt.Println(Area(r)) // 12
}
```

---

## Lesson 5.2 — Pointers (`&` / `*`)

### Syntax
```go
p := &x   // "address of" — p is a *T pointing at x
*p        // "dereference" — read the value p points to
*p = 10   // write THROUGH the pointer, mutating the original x
```

A **pointer** holds the *memory address* of a variable instead of a copy of its value. You create one with `&` ("address of") and read/write through it with `*` ("dereference").

```go
x := 42
p := &x    // p is a *int — a pointer to x
fmt.Println(p)  // 0xc0000b4008 (some address)
fmt.Println(*p) // 42 — dereference
```

The power of a pointer is that writing through it modifies the *original* variable:

```go
x := 7
p := &x
*p = 10 // writes through the pointer
fmt.Println(x) // 10 — x changed!
```

> **Key point:** A function that receives a pointer can mutate the caller's variable. A function that receives a plain value only gets a copy. This distinction is what makes **pointer receivers** in Go methods work — we'll build on it in Lesson 5.3.

```go
package main

import "fmt"

func main() {
    x := 7
    p := &x
    *p = 10
    fmt.Println(x) // 10
}
```

---

## Lesson 5.3 — Methods & Pointer Receivers

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

---

## Lesson 5.4 — Interfaces

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

---

## Lesson 5.5 — Struct Embedding

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

---

## Lesson 5.6 — Capstone: Café Menu

This lesson introduces **no new concept** — it's a consolidation drill that combines structs, methods, interfaces, and embedding from Lessons 5.1–5.5 into one tiny café menu app.

```go
type Product struct {
    Name      string
    BasePrice float64
}

func (p Product) Label() string {
    return fmt.Sprintf("%s: $%.2f", p.Name, p.BasePrice)
}
```

The pattern you'll keep using: a struct holds the data, methods attach behavior, an interface lets different kinds of products be treated uniformly, and embedding lets you extend a product without re-declaring its fields.

> **Reminder:** `Sprintf` builds a string, it doesn't print. Combine it with the `%s` and `%.2f` verbs you learned in Chapter 1. Also note the stored amount is a field named `BasePrice` — a Go struct cannot have both a field and a method with the same name.

```go
package main

import "fmt"

type Product struct {
    Name      string
    BasePrice float64
}

func (p Product) Label() string {
    return fmt.Sprintf("%s: $%.2f", p.Name, p.BasePrice)
}

func main() {
    p := Product{Name: "Coffee", BasePrice: 4.50}
    fmt.Println(p.Label()) // Coffee: $4.50
}
```