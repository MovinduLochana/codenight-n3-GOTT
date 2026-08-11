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