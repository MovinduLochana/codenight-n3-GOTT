# Chapter 4 — Functions
## Task Slides (Practice Screens)

Each task below is the split-screen practice slide shown right after the matching content slide.
**LHS** = task description shown to the student. **RHS** = starter code pre-loaded in the editor.
`// TODO` marks exactly where the student needs to write code.

---

## Task 4.1a — `Divide` (Beginner)

**LHS — Task Description**
> Complete `Divide(a, b int) (int, error)`. Return an error (using `errors.New`) if `b` is `0`; otherwise return `a / b` and `nil`.
>
> **Expected behavior:**
> ```go
> Divide(10, 2) // 5, nil
> Divide(5, 0)  // 0, error
> ```

**RHS — Starter Code**
```go
package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	// TODO: return an error if b == 0, otherwise return a / b and nil
	return 0, nil
}

func main() {
	result, err := Divide(10, 2)
	fmt.Println(result, err)

	result, err = Divide(5, 0)
	fmt.Println(result, err)
}
```

---

## Task 4.1b — `DivMod` (Intermediate)

**LHS — Task Description**
> Complete `DivMod(a, b int) (quotient, remainder int, err error)`. It should return three values: the quotient and remainder of `a / b` (integer division and `%`), plus an error if `b == 0`. Use Go's **named return values** so you can return them individually (`return 0, 0, errors.New(...)`).
>
> **Expected behavior:**
> ```go
> q, r, err := DivMod(10, 3) // 3, 1, nil
> q, r, err := DivMod(10, 0) // 0, 0, error
> ```

**RHS — Starter Code**
```go
package main

import (
	"errors"
	"fmt"
)

func DivMod(a, b int) (quotient, remainder int, err error) {
	// TODO: return quotient = a/b, remainder = a%b, err = nil when b != 0;
	// return 0, 0, errors.New(...) when b == 0
	return 0, 0, nil
}

func main() {
	q, r, err := DivMod(10, 3)
	fmt.Println(q, r, err)

	q, r, err = DivMod(10, 0)
	fmt.Println(q, r, err)
}
```

---

## Task 4.1c — `MinMax` (Advanced)

**LHS — Task Description**
> Complete `MinMax(nums []int) (min, max int)`. It should find and return both the smallest and largest element of `nums` in one call. Use a `range` loop and the multiple-return form. You can assume `nums` is non-empty.
>
> **Expected behavior:**
> ```go
> MinMax([]int{3, 1, 4, 1, 5}) // 1, 5
> MinMax([]int{7})             // 7, 7
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func MinMax(nums []int) (min, max int) {
	// TODO: find and return the smallest and largest element in one pass
	return 0, 0
}

func main() {
	lo, hi := MinMax([]int{3, 1, 4, 1, 5})
	fmt.Println(lo, hi) // 1 5

	lo, hi = MinMax([]int{7})
	fmt.Println(lo, hi) // 7 7
}
```

---

## Task 4.2a — `Max` (Beginner)

**LHS — Task Description**
> Complete `Max(nums ...int) int` using a variadic parameter. Return the largest value passed in. If no arguments are given, return `0`.
>
> **Expected behavior:**
> ```go
> Max(3, 7, 2) // 7
> Max(5)       // 5
> Max()        // 0
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Max(nums ...int) int {
	// TODO: return the largest value in nums, or 0 if nums is empty
	return 0
}

func main() {
	fmt.Println(Max(3, 7, 2))
	fmt.Println(Max(5))
	fmt.Println(Max())
}
```

---

## Task 4.2b — `Join` (Intermediate)

**LHS — Task Description**
> Complete `Join(sep string, parts ...string) string`. It should concatenate all `parts` with `sep` between them, using a `range` loop and the `+` operator. It should never place a separator before the first or after the last part.
>
> **Expected behavior:**
> ```go
> Join(", ", "go", "is", "fun") // "go, is, fun"
> Join("|", "a")                // "a"
> Join("-")                     // ""
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Join(sep string, parts ...string) string {
	// TODO: concatenate parts with sep between them (never at the ends)
	return ""
}

func main() {
	fmt.Println(Join(", ", "go", "is", "fun"))
	fmt.Println(Join("|", "a"))
	fmt.Println(Join("-"))
}
```

---

## Task 4.2c — `Mean` (Advanced)

**LHS — Task Description**
> Complete `Mean(scores ...float64) float64`. It should return the arithmetic average of all scores. If no scores are given, return `0`. In `main`, also call `Mean` by **spreading a slice** with `values...` to prove the two forms are equivalent.
>
> **Expected behavior:**
> ```go
> Mean(90, 86, 88)    // 88
> Mean(10, 20)        // 15
> Mean()              // 0
> values := []float64{2, 4} ; Mean(values...) // 3
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Mean(scores ...float64) float64 {
	// TODO: return the average of all scores, or 0 if empty
	return 0
}

func main() {
	fmt.Println(Mean(90, 86, 88))
	fmt.Println(Mean(10, 20))
	fmt.Println(Mean())

	values := []float64{2, 4}
	fmt.Println(Mean(values...)) // spread a slice into the variadic call
}
```

---

## Task 4.3a — `MakeCounter` (Beginner)

**LHS — Task Description**
> Complete `MakeCounter() func() int`. Each time the returned function is called, it should return an incrementing count, starting at `1`.
>
> **Expected behavior:**
> ```go
> counter := MakeCounter()
> counter() // 1
> counter() // 2
> counter() // 3
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func MakeCounter() func() int {
	// TODO: return a closure that increments and returns a count, starting at 1
	return nil
}

func main() {
	counter := MakeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
```

---

## Task 4.3b — `MakeAdder` (Intermediate)

**LHS — Task Description**
> Complete `MakeAdder(add int) func(int) int`. The returned function should **capture** `add` (a closure) and add it to whatever argument it receives.
>
> **Expected behavior:**
> ```go
> plus5 := MakeAdder(5)
> plus5(10) // 15
> plus5(0)  // 5
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func MakeAdder(add int) func(int) int {
	// TODO: return a closure that captures `add` and adds it to its argument
	return nil
}

func main() {
	plus5 := MakeAdder(5)
	fmt.Println(plus5(10)) // 15
	fmt.Println(plus5(0))  // 5
}
```

---

## Task 4.3c — `MakeBank` (Advanced)

**LHS — Task Description**
> Complete `MakeBank(balance float64) func(amount float64) float64`. The returned closure captures `balance` and every call **deposits** `amount` (if positive) and returns the new running balance. If `amount` is negative or zero, leave the balance unchanged and return the current balance. Two independent banks must never share state.
>
> **Expected behavior:**
> ```go
> bank := MakeBank(100)
> bank(50)  // 150
> bank(-20) // 150 (ignored)
> bank(25)  // 175
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func MakeBank(balance float64) func(amount float64) float64 {
	// TODO: return a closure capturing `balance` that:
	//   - deposits positive amounts and returns the new balance
	//   - ignores zero/negative amounts (returns current balance unchanged)
	return nil
}

func main() {
	bank := MakeBank(100)
	fmt.Println(bank(50))  // 150
	fmt.Println(bank(-20)) // 150
	fmt.Println(bank(25))  // 175
}
```

---

## Task 4.4a — `ProcessLog` (Beginner)

**LHS — Task Description**
> Complete `ProcessLog() (log []string)`. Append `"start"` and `"middle"` normally, then use exactly **one** `defer` to append `"end"` **after** them, using a **named return value** and a bare `return`.
>
> > **Key point:** the deferred append must go through the *named* return value (`log`). If you used `return log` with an anonymous return, the return value would be snapshotted *before* the deferred call runs and `"end"` would be lost.
>
> **Expected behavior:**
> ```go
> ProcessLog() // ["start", "middle", "end"]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func ProcessLog() (log []string) {
	// TODO: append "start" and "middle" normally

	// TODO: use defer to append "end" through the named return value `log`
	// (so it runs after the normal appends)

	return
}

func main() {
	fmt.Println(ProcessLog())
}
```

---

## Task 4.4b — `StackLog` (Intermediate)

**LHS — Task Description**
> Complete `StackLog() (log []string)`. Schedule **two** deferred appends on the named return: defer `"B"` **first**, then defer `"A"`. Because defers run **LIFO**, `"A"` (deferred last) must be appended before `"B"`. Append `"start"` normally first. Expected `["start", "A", "B"]`.
>
> **Expected behavior:**
> ```go
> StackLog() // ["start", "A", "B"]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func StackLog() (log []string) {
	// TODO: append "start" normally

	// TODO: defer appending "B", THEN defer appending "A".
	// LIFO means "A" (deferred last) runs first.
	// Expected final log: ["start", "A", "B"]

	return
}

func main() {
	fmt.Println(StackLog())
}
```

---

## Task 4.4c — `LoopCleanup` (Advanced)

**LHS — Task Description**
> Complete `LoopCleanup(n int) (nums []int)`. Inside a `for i := 0; i < n; i++` loop, defer an append of `i` to the named return. Because deferred calls run **after** the loop, in LIFO order, the collected numbers come back **reversed**: `n-1` down to `0`.
>
> **Expected behavior (Go 1.22+ per-iteration loop semantics):**
> ```go
> LoopCleanup(3) // [2 1 0]
> LoopCleanup(1) // [0]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func LoopCleanup(n int) (nums []int) {
	// TODO: for i := 0; i < n; i++ {
	//   defer appending i to the named return value `nums`
	// }
	// Deferred calls run after the loop in LIFO order,
	// so the result comes back reversed: n-1 ... 0
	return
}

func main() {
	fmt.Println(LoopCleanup(3)) // [2 1 0]
	fmt.Println(LoopCleanup(1)) // [0]
}
```

---

## Quick Reference — What Each Task's Hidden Test Checks

| Task | Test checks |
|------|-------------|
| 4.1a | `Divide(10, 2)` returns `5, nil`; `Divide(5, 0)` returns error |
| 4.1b | `DivMod(10, 3)` returns `3, 1, nil`; `DivMod(10, 0)` returns error |
| 4.1c | `MinMax(...)` returns `(1, 5)` and `(7, 7)` |
| 4.2a | `Max(3, 7, 2) == 7`, `Max(5) == 5`, `Max() == 0` |
| 4.2b | `Join(", ", ...) == "go, is, fun"`, `Join("|", "a") == "a"`, `Join("-") == ""` |
| 4.2c | `Mean(90, 86, 88) == 88`, `Mean() == 0`, slice-spread form equals value form |
| 4.3a | successive calls return `1, 2, 3` |
| 4.3b | `plus5(10) == 15`, `plus5(0) == 5` |
| 4.3c | `bank(50) == 150`, `bank(-20) == 150`, `bank(25) == 175`; independent balances |
| 4.4a | `ProcessLog() == ["start", "middle", "end"]` |
| 4.4b | `StackLog() == ["start", "A", "B"]` |
| 4.4c | `LoopCleanup(3) == [2 1 0]`, `LoopCleanup(1) == [0]` |