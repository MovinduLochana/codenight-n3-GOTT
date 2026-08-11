# Chapter 1 — Go Fundamentals
## Task Slides (Practice Screens)

Each task below is the split-screen practice slide shown right after the matching content slide.
**LHS** = task description shown to the student. **RHS** = starter code pre-loaded in the editor.
`// TODO` marks exactly where the student needs to write code.

---

## Task 1.1 — Declaring Variables with `var`

**LHS — Task Description**
> Declare a variable named `city` of type `string` using the `var` keyword, and assign it the value `"Colombo"`. Then print it using `fmt.Println`.
>
> **Expected Output:**
> ```
> Colombo
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func main() {
	// TODO: declare a variable named `city` of type string using `var`,
	// and assign it the value "Colombo"

	fmt.Println(city)
}
```

---

## Task 1.2 — Short Declaration `:=`

**LHS — Task Description**
> Declare the same `city` variable as before, but this time use the short declaration operator `:=` instead of `var`. Assign it the value `"Colombo"` and print it.
>
> **Expected Output:**
> ```
> Colombo
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func main() {
	// TODO: declare `city` using := and assign it "Colombo"

	fmt.Println(city)
}
```

---

## Task 1.3 — Zero Values

**LHS — Task Description**
> Declare two variables without assigning them a value:
> - `score` of type `int`
> - `passed` of type `bool`
>
> Print both on one line using `fmt.Println` and observe their zero values.
>
> **Expected Output:**
> ```
> 0 false
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func main() {
	// TODO: declare `score int` and `passed bool` with no assigned value

	fmt.Println(score, passed)
}
```

---

## Task 1.4 — Basic Types Tour

**LHS — Task Description**
> Declare one variable of each of the four basic types, using any values you like:
> - `age` — `int`
> - `price` — `float64`
> - `city` — `string`
> - `isOpen` — `bool`
>
> Print all four on a single line, space-separated, using `fmt.Println`.
>
> **Expected Output (example — your values may differ):**
> ```
> 30 19.99 Colombo true
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func main() {
	// TODO: declare `age` (int), `price` (float64), `city` (string), `isOpen` (bool)

	fmt.Println(age, price, city, isOpen)
}
```

---

## Task 1.5 — Constants (`const`)

**LHS — Task Description**
> Declare a constant named `MaxScore` with the value `100` using the `const` keyword, then print it.
>
> **Expected Output:**
> ```
> 100
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func main() {
	// TODO: declare `const MaxScore = 100`

	fmt.Println(MaxScore)
}
```

---

## Task 1.6 — `iota` for Enums

**LHS — Task Description**
> Using an `iota`-based `const` block, define three constants in this exact order: `Low`, `Medium`, `High`.
>
> Then complete the function `PriorityValue()` so it returns the numeric value of `Medium`.
>
> **Expected behavior:**
> ```go
> PriorityValue() // returns 1
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

// TODO: define a const block using iota with Low, Medium, High in that order

func PriorityValue() int {
	// TODO: return the value of Medium
	return 0
}

func main() {
	fmt.Println(PriorityValue())
}
```

---

## Task 1.7 — Operators

**LHS — Task Description**
> Complete the function `IsEven(n int) bool` so it returns `true` if `n` is even and `false` otherwise. Use the `%` (modulo) operator.
>
> **Expected behavior:**
> ```go
> IsEven(4) // true
> IsEven(7) // false
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func IsEven(n int) bool {
	// TODO: return true if n is even, false otherwise
	return false
}

func main() {
	fmt.Println(IsEven(4)) // should print true
	fmt.Println(IsEven(7)) // should print false
}
```

---

## Task 1.8 — Type Conversion

**LHS — Task Description**
> Complete the function `Average(a, b int) float64` so it returns the true decimal average of `a` and `b` — not a truncated integer result. You'll need to convert `a` and `b` to `float64` before dividing.
>
> **Expected behavior:**
> ```go
> Average(3, 4) // 3.5
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Average(a, b int) float64 {
	// TODO: convert a and b to float64 before dividing, then return the result
	return 0
}

func main() {
	fmt.Println(Average(3, 4)) // should print 3.5
}
```

---

## Task 1.9 — `fmt` Print Verbs

**LHS — Task Description**
> Given the variable `age := 30`, use `fmt.Printf` to print its value and its type in this exact format, using the `%d` and `%T` verbs:
>
> **Expected Output:**
> ```
> Age: 30, Type: int
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func main() {
	age := 30

	// TODO: use fmt.Printf with %d and %T to print "Age: 30, Type: int"
}
```

---

## Task 1.10 — String Formatting Practice

**LHS — Task Description**
> Complete the function `Receipt(item string, qty int, price float64) string`. It should return a string built with `fmt.Sprintf` in this exact format:
> ```
> Item: <item>, Qty: <qty>, Price: $<price to 2 decimal places>
> ```
>
> **Expected behavior:**
> ```go
> Receipt("Coffee", 2, 4.50) // "Item: Coffee, Qty: 2, Price: $4.50"
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Receipt(item string, qty int, price float64) string {
	// TODO: use fmt.Sprintf to build and return the formatted receipt string
	return ""
}

func main() {
	fmt.Println(Receipt("Coffee", 2, 4.50))
}
```

---

## Quick Reference — What Each Task's Hidden Test Checks

| Task | Test checks |
|------|-------------|
| 1.1 | stdout == `Colombo` |
| 1.2 | stdout == `Colombo` |
| 1.3 | stdout == `0 false` |
| 1.4 | stdout has 4 space-separated tokens matching int / float / string / bool shapes |
| 1.5 | stdout == `100` |
| 1.6 | `PriorityValue() == 1` |
| 1.7 | `IsEven(4) == true`, `IsEven(7) == false`, plus 2–3 extra cases |
| 1.8 | `Average(3, 4) == 3.5`, plus extra cases |
| 1.9 | stdout == `Age: 30, Type: int` |
| 1.10 | `Receipt(...)` returns exact expected string for 2–3 sample inputs |
