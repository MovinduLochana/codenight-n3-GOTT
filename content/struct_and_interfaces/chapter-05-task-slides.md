# Chapter 5 — Structs & Interfaces
## Task Slides (Practice Screens)

Each task below is the split-screen practice slide shown right after the matching content slide.
**LHS** = task description shown to the student. **RHS** = starter code pre-loaded in the editor.
`// TODO` marks exactly where the student needs to write code.

---

## Task 5.1a — `Rectangle` + `Area` (Beginner)

**LHS — Task Description**
> Define a struct `Rectangle` with fields `Width` and `Height` (both `float64`). Then complete `Area(r Rectangle) float64` to return `Width * Height`.
>
> **Expected behavior:**
> ```go
> Area(Rectangle{Width: 3, Height: 4}) // 12
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

// TODO: define the Rectangle struct with Width and Height (float64)

func Area(r Rectangle) float64 {
	// TODO: return r.Width * r.Height
	return 0
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println(Area(r))
}
```

---

## Task 5.1b — `Box` + `Volume` (Intermediate)

**LHS — Task Description**
> Define a struct `Box` with fields `Length`, `Width`, and `Height` (all `float64`). Then complete `Volume(b Box) float64` to return `Length * Width * Height`.
>
> **Expected behavior:**
> ```go
> Volume(Box{Length: 2, Width: 3, Height: 4}) // 24
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

// TODO: define the Box struct with Length, Width, Height (all float64)

func Volume(b Box) float64 {
	// TODO: return b.Length * b.Width * b.Height
	return 0
}

func main() {
	fmt.Println(Volume(Box{Length: 2, Width: 3, Height: 4}))
}
```

---

## Task 5.1c — `Product` + `Cart` + `AddItem` (Advanced)

**LHS — Task Description**
> Define a struct `Product` with fields `Name string` and `Price float64`, and a struct `Cart` with a slice field `Items []Product`. Then complete `AddItem(c Cart, p Product) Cart` — it should return a **new** `Cart` with `p` appended to `c.Items` (copy the slice header; no pointers needed yet).
>
> **Expected behavior:**
> ```go
> c := Cart{}
> c = AddItem(c, Product{Name: "Coffee", Price: 4.50})
> c = AddItem(c, Product{Name: "Tea", Price: 3.00})
> len(c.Items)   // 2
> c.Items[0].Name // "Coffee"
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

// TODO: define the Product struct (Name string, Price float64)
// TODO: define the Cart struct with field Items []Product

func AddItem(c Cart, p Product) Cart {
	// TODO: return a new Cart with p appended to c.Items
	return c
}

func main() {
	c := Cart{}
	c = AddItem(c, Product{Name: "Coffee", Price: 4.50})
	c = AddItem(c, Product{Name: "Tea", Price: 3.00})
	fmt.Println(len(c.Items), c.Items[0].Name)
}
```

---

## Task 5.2a — `SetFive` (Beginner)

**LHS — Task Description**
> Complete `SetFive(ptr *int)` so it sets the variable pointed to by `ptr` to `5` (write through the pointer with `*ptr = 5`).
>
> **Expected behavior:**
> ```go
> x := 3
> SetFive(&x)
> x // 5 — the original variable changed
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func SetFive(ptr *int) {
	// TODO: set the variable pointed to by ptr to 5
}

func main() {
	x := 3
	SetFive(&x)
	fmt.Println(x) // 5
}
```

---

## Task 5.2b — `Double` (Intermediate)

**LHS — Task Description**
> Complete `Double(ptr *int) int`. It should multiply the value pointed to by `ptr` by 2 (writing back through the pointer), and return the new value.
>
> **Expected behavior:**
> ```go
> x := 7
> Double(&x) // 14
> x          // 14 — mutated through the pointer
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Double(ptr *int) int {
	// TODO: double the value pointed to by ptr, write it back, and return it
	return 0
}

func main() {
	x := 7
	fmt.Println(Double(&x)) // 14
	fmt.Println(x)          // 14
}
```

---

## Task 5.2c — `SwapNums` (Advanced)

**LHS — Task Description**
> Complete `SwapNums(a, b *int)` so it exchanges the values pointed to by `a` and `b` — no return value needed. This is the classic "how would you do this without pointers?" problem.
>
> **Expected behavior:**
> ```go
> x, y := 1, 2
> SwapNums(&x, &y)
> x // 2
> y // 1
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func SwapNums(a, b *int) {
	// TODO: exchange the values pointed to by a and b
	// use a temporary variable
}

func main() {
	x, y := 1, 2
	SwapNums(&x, &y)
	fmt.Println(x, y) // 2 1
}
```

---

## Task 5.3a — `Scale` (Beginner)

**LHS — Task Description**
> Add a method `Scale(factor float64)` on `*Rectangle` (pointer receiver) that multiplies both `Width` and `Height` by `factor`, mutating the original struct.
>
> **Expected behavior:**
> ```go
> r := Rectangle{Width: 3, Height: 4}
> r.Scale(2)
> // r.Width == 6, r.Height == 8
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// TODO: define Scale(factor float64) on *Rectangle

func main() {
	r := Rectangle{Width: 3, Height: 4}
	r.Scale(2)
	fmt.Println(r.Width, r.Height)
}
```

---

## Task 5.3b — `Money` with `Format` (Intermediate)

**LHS — Task Description**
> Define a type `Money float64` with a **value receiver** method `Format()` that returns the amount as a `float64` (just returns `float64(m)`). Then define `func AddCurrency(a, b Money) Money` that adds two `Money` values. The point: `Money` is a named type with methods, not a struct.
>
> **Expected behavior:**
> ```go
> m := Money(12.5)
> m.Format() // 12.5
> AddCurrency(Money(1.25), Money(2.50)) // Money(3.75)
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

// TODO: define `type Money float64`

// TODO: define a value-receiver method Format() float64 on Money
// that returns float64(m)

// TODO: define AddCurrency(a, b Money) Money

func main() {
	m := Money(12.5)
	fmt.Println(m.Format())                          // 12.5
	fmt.Println(AddCurrency(Money(1.25), Money(2.50))) // 3.75
}
```

---

## Task 5.3c — `Stack` (Advanced)

**LHS — Task Description**
> Define a type `Stack` with a field `Items []int`. Add **pointer-receiver** methods that mutate the stack:
> - `Push(v int)` — appends `v` to the stack
> - `Pop() int` — removes and returns the **last** item
>
> A **pointer receiver** is required because both methods mutate the receiver (including `Pop`, which must shrink the slice and return the removed value). You can assume `Pop` is only called on a non-empty stack.
>
> **Expected behavior:**
> ```go
> s := Stack{}
> s.Push(1); s.Push(2); s.Push(3)
> s.Pop() // 3
> s.Pop() // 2
> s.Items // [1]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

// TODO: define `type Stack struct { Items []int }`

// TODO: define Push(v int) on *Stack — append v to Items

// TODO: define Pop() int on *Stack — remove and return the last item

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop()) // 3
	fmt.Println(s.Pop()) // 2
	fmt.Println(s.Items) // [1]
}
```

---

## Task 5.4a — `Circle` + `TotalArea` (Beginner)

**LHS — Task Description**
> Define a struct `Circle` with field `Radius float64` and give it an `Area() float64` method (use `3.14159` for pi). Then complete `TotalArea(shapes []Shape) float64` to sum every shape's area. `Circle` automatically satisfies the `Shape` interface — no `implements` declaration needed.
>
> **Expected behavior:**
> ```go
> shapes := []Shape{
>     Rectangle{Width: 2, Height: 3}, // 6
>     Circle{Radius: 1},              // ≈ 3.14
> }
> TotalArea(shapes) // ≈ 9.14
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TODO: define Circle struct with field Radius float64
// TODO: give Circle an Area() float64 method (use 3.14159 for pi)

func TotalArea(shapes []Shape) float64 {
	// TODO: sum the Area() of every shape in the slice
	return 0
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
	}
	fmt.Println(TotalArea(shapes))
}
```

---

## Task 5.4b — `Perimeterer` + `TotalPerimeter` (Intermediate)

**LHS — Task Description**
> Define an interface `Perimeterer` with one method `Perimeter() float64`. Give `Rectangle` a `Perimeter()` method returning `2 * (Width + Height)`, and `Circle` a `Perimeter()` returning `2 * 3.14159 * Radius`. Then complete `TotalPerimeter(shapes []Perimeterer) float64` to sum every shape's perimeter. This is behavior polymorphism: two different types, one interface.
>
> **Expected behavior:**
> ```go
> shapes := []Perimeterer{
>     Rectangle{Width: 2, Height: 3}, // 10
>     Circle{Radius: 1},              // ≈ 6.28
> }
> TotalPerimeter(shapes) // ≈ 16.28
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

// TODO: define interface Perimeterer with method Perimeter() float64

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Perimeter() float64 {
	// TODO: return 2 * (Width + Height)
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	// TODO: return 2 * 3.14159 * Radius
	return 0
}

func TotalPerimeter(shapes []Perimeterer) float64 {
	// TODO: sum every shape's Perimeter()
	return 0
}

func main() {
	shapes := []Perimeterer{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
	}
	fmt.Println(TotalPerimeter(shapes))
}
```

---

## Task 5.4c — `Biggest` (Advanced)

**LHS — Task Description**
> Complete `Biggest(shapes []Shape) float64`. It should return the **largest** `Area()` among the shapes in the slice. Use a `range` loop over `[]Shape` tracking the max seen so far. (Assumes at least one shape.)
>
> **Expected behavior:**
> ```go
> shapes := []Shape{
>     Rectangle{Width: 2, Height: 3}, // 6
>     Circle{Radius: 2},              // ≈ 12.57
> }
> Biggest(shapes) // ≈ 12.57
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func Biggest(shapes []Shape) float64 {
	// TODO: return the largest Area() among the shapes
	return 0
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 2},
	}
	fmt.Println(Biggest(shapes))
}
```

---

## Task 5.5a — `NamedRectangle` (Beginner)

**LHS — Task Description**
> Define `NamedRectangle` that **embeds** `Rectangle` (no field name) and adds a `Name string` field. Because `Rectangle` is embedded, `NamedRectangle` automatically gets the promoted `Area()` method and `Width`/`Height` fields.
>
> **Expected behavior:**
> ```go
> nr := NamedRectangle{Rectangle: Rectangle{Width: 3, Height: 4}, Name: "MyRect"}
> nr.Name  // "MyRect"
> nr.Area() // 12 — promoted from Rectangle
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TODO: define NamedRectangle, embedding Rectangle, plus a Name string field

func main() {
	nr := NamedRectangle{Rectangle: Rectangle{Width: 3, Height: 4}, Name: "MyRect"}
	fmt.Println(nr.Name, nr.Area())
}
```

---

## Task 5.5b — `WeightedRect` (Intermediate)

**LHS — Task Description**
> Define `WeightedRect` that **embeds** `Rectangle` and adds a `Weight float64` field. Give `WeightedRect` its **own** `Area() float64` method that returns the promoted area scaled by the weight: `wr.Rectangle.Area() * wr.Weight`. This shadows the promoted `Rectangle.Area()` for `WeightedRect` values.
>
> **Expected behavior:**
> ```go
> wr := WeightedRect{Rectangle: Rectangle{Width: 3, Height: 4}, Weight: 2}
> wr.Area()          // 24 (own method wins)
> wr.Rectangle.Area() // 12
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TODO: define WeightedRect embedding Rectangle, plus a Weight float64 field
// TODO: give WeightedRect its OWN Area() method that returns
//   wr.Rectangle.Area() * wr.Weight

func main() {
	wr := WeightedRect{Rectangle: Rectangle{Width: 3, Height: 4}, Weight: 2}
	fmt.Println(wr.Area())           // 24
	fmt.Println(wr.Rectangle.Area()) // 12
}
```

---

## Task 5.5c — `Square` satisfies `Shape` (Advanced)

**LHS — Task Description**
> Define `Square` that embeds `Rectangle` and adds a `Side float64` field. Give `Square` its **own** `Area() float64` method returning `Side * Side`, shadowing the promoted `Rectangle.Area()`. Then `Square` satisfies the `Shape` interface through its own method — complete `TotalArea(shapes []Shape)`.
>
> **Expected behavior:**
> ```go
> shapes := []Shape{
>     Square{Rectangle: Rectangle{Width: 3, Height: 4}, Side: 5},
> }
> TotalArea(shapes) // 25 (Square.Area() = Side*Side, not 3*4)
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TODO: define Square embedding Rectangle, plus a Side float64 field
// TODO: give Square its OWN Area() method returning Side * Side,
//   shadowing the promoted Rectangle.Area()

func TotalArea(shapes []Shape) float64 {
	// TODO: sum the Area() of every shape
	return 0
}

func main() {
	shapes := []Shape{
		Square{Rectangle: Rectangle{Width: 3, Height: 4}, Side: 5},
	}
	fmt.Println(TotalArea(shapes))
}
```

---

## Task 5.6a — `Product.Label` (Beginner)

**LHS — Task Description**
> Define a struct `Product` with fields `Name string` and `BasePrice float64`. Give `Product` a `Label() string` method that returns `fmt.Sprintf("%s: $%.2f", p.Name, p.BasePrice)` — the `%s`/`%.2f` verbs from Chapter 1. (The stored amount is a field called `BasePrice` because a struct can't have both a field and a method named `Price`.)
>
> **Expected behavior:**
> ```go
> Product{Name: "Coffee", BasePrice: 4.50}.Label() // "Coffee: $4.50"
> Product{Name: "Tea", BasePrice: 3.00}.Label()    // "Tea: $3.00"
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

// TODO: define Product struct with Name string and BasePrice float64
// TODO: give Product a Label() string method returning
//   fmt.Sprintf("%s: $%.2f", p.Name, p.BasePrice)

func main() {
	fmt.Println(Product{Name: "Coffee", BasePrice: 4.50}.Label())
	fmt.Println(Product{Name: "Tea", BasePrice: 3.00}.Label())
}
```

---

## Task 5.6b — `Pricer` + `TotalPrice` (Intermediate)

**LHS — Task Description**
> Define an interface `Pricer` with one method `Price() float64`. Give `Product` a `Price() float64` method that returns `p.BasePrice`. Then complete `TotalPrice(items []Pricer) float64` which sums every item's price. This is "many products, one interface" — the same pattern as `Shape`/`Area()`.
>
> **Expected behavior:**
> ```go
> items := []Pricer{
>     Product{Name: "Coffee", BasePrice: 4.50},
>     Product{Name: "Cake", BasePrice: 6.00},
> }
> TotalPrice(items) // 10.5
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

// TODO: define interface Pricer with method Price() float64

type Product struct {
	Name      string
	BasePrice float64
}

func (p Product) Price() float64 {
	// TODO: return p.BasePrice
	return 0
}

func TotalPrice(items []Pricer) float64 {
	// TODO: sum every item's Price()
	return 0
}

func main() {
	items := []Pricer{
		Product{Name: "Coffee", BasePrice: 4.50},
		Product{Name: "Cake", BasePrice: 6.00},
	}
	fmt.Println(TotalPrice(items))
}
```

---

## Task 5.6c — `Deal` (Advanced)

**LHS — Task Description**
> Define `Deal` that **embeds** `Product` and adds a `Discount float64` field. Give `Deal` its **own** `Price() float64` method that returns the discounted price: `d.Product.BasePrice * (1 - d.Discount)` (so a `0.20` discount means 20% off). Because `Deal` has its own `Price()` method, it satisfies `Pricer` polymorphically — the same `TotalPrice([]Pricer)` from the previous task keeps working.
>
> **Expected behavior:**
> ```go
> items := []Pricer{
>     Product{Name: "Coffee", BasePrice: 4.00},
>     Deal{Product: Product{Name: "Cake", BasePrice: 10.00}, Discount: 0.20},
> }
> TotalPrice(items) // 4.00 + 8.00 = 12.0
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

type Pricer interface {
	Price() float64
}

type Product struct {
	Name      string
	BasePrice float64
}

func (p Product) Price() float64 {
	return p.BasePrice
}

// TODO: define Deal embedding Product, plus a Discount float64 field
// TODO: give Deal its OWN Price() float64 method returning
//   d.Product.BasePrice * (1 - d.Discount)

func TotalPrice(items []Pricer) float64 {
	// TODO: sum every item's Price()
	return 0
}

func main() {
	items := []Pricer{
		Product{Name: "Coffee", BasePrice: 4.00},
		Deal{Product: Product{Name: "Cake", BasePrice: 10.00}, Discount: 0.20},
	}
	fmt.Println(TotalPrice(items))
}
```

---

## Quick Reference — What Each Task's Hidden Test Checks

| Task | Test checks |
|------|-------------|
| 5.1a | `Area(Rectangle{3, 4}) == 12` |
| 5.1b | `Volume(Box{2, 3, 4}) == 24` |
| 5.1c | after two `AddItem`, `len(Items) == 2` and `Items[0].Name == "Coffee"` |
| 5.2a | `x == 5` after `SetFive(&x)` |
| 5.2b | `Double(&x) == 14` and `x == 14` |
| 5.2c | after `SwapNums(&x,&y)`, `x == 2 && y == 1` |
| 5.3a | `r.Width == 6` and `r.Height == 8` after `Scale(2)` |
| 5.3b | `Format() == 12.5`; `AddCurrency(1.25, 2.50) == 3.75` |
| 5.3c | `Pop()` returns `3` then `2`; `s.Items == [1]` |
| 5.4a | `TotalArea(...) ≈ 9.14` |
| 5.4b | `TotalPerimeter(...) ≈ 16.28` |
| 5.4c | `Biggest(...) ≈ 12.57` |
| 5.5a | `nr.Name == "MyRect"` and `nr.Area() == 12` |
| 5.5b | `wr.Area() == 24` and `wr.Rectangle.Area() == 12` |
| 5.5c | `TotalArea(...) == 25` |
| 5.6a | `Label()` returns exact `"Coffee: $4.50"` / `"Tea: $3.00"` strings |
| 5.6b | `TotalPrice(...) == 10.5` |
| 5.6c | `TotalPrice(...) == 12.0` |