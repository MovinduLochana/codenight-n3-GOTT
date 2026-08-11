# Chapter 2 — Control Flow
## Task Slides (Practice Screens)

Each task below is the split-screen practice slide shown right after the matching content slide.
**LHS** = task description shown to the student. **RHS** = starter code pre-loaded in the editor.
`// TODO` marks exactly where the student needs to write code.

---

## Task 2.1a — `Grade` (Beginner)

**LHS — Task Description**
> Complete `Grade(score int) string`. Return:
> - `"A"` for 90 and above
> - `"B"` for 80–89
> - `"C"` for 70–79
> - `"F"` for anything below 70
>
> **Expected behavior:**
> ```go
> Grade(95) // "A"
> Grade(82) // "B"
> Grade(71) // "C"
> Grade(50) // "F"
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Grade(score int) string {
	// TODO: return "A", "B", "C", or "F" based on score
	return ""
}

func main() {
	fmt.Println(Grade(95))
	fmt.Println(Grade(82))
	fmt.Println(Grade(71))
	fmt.Println(Grade(50))
}
```

---

## Task 2.1b — `Sign` (Intermediate)

**LHS — Task Description**
> Complete `Sign(n int) string`. Return `"positive"` if `n > 0`, `"negative"` if `n < 0`, and `"zero"` if `n == 0`. Use an `if` / `else if` / `else` chain.
>
> **Expected behavior:**
> ```go
> Sign(42)  // "positive"
> Sign(-3)  // "negative"
> Sign(0)   // "zero"
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Sign(n int) string {
	// TODO: return "positive", "negative", or "zero"
	return ""
}

func main() {
	fmt.Println(Sign(42))
	fmt.Println(Sign(-3))
	fmt.Println(Sign(0))
}
```

---

## Task 2.1c — `ClassFare` (Advanced)

**LHS — Task Description**
> A café charges a flat fare. Complete `ClassFare(age int, isStudent bool) int`:
> - children under 12 and seniors 65+ pay `25`
> - students pay `30` (unless they already qualify for the `25` fare)
> - everyone else pays `50`
>
> Use an `if` / `else if` chain with `&&`/`||`, and check the special fares first.
>
> **Expected behavior:**
> ```go
> ClassFare(8, false)   // 25
> ClassFare(70, false)  // 25
> ClassFare(20, true)   // 30
> ClassFare(20, false)  // 50
> ClassFare(10, true)   // 25 (age < 12 qualifies first — special fare wins)
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func ClassFare(age int, isStudent bool) int {
	// TODO: return the fare based on age and student status
	//   under 12 or 65+ → 25
	//   student         → 30 (unless already 25)
	//   everyone else   → 50
	return 0
}

func main() {
	fmt.Println(ClassFare(8, false))
	fmt.Println(ClassFare(70, false))
	fmt.Println(ClassFare(20, true))
	fmt.Println(ClassFare(20, false))
	fmt.Println(ClassFare(10, true))
}
```

---

## Task 2.2a — `SumEvens` (Beginner)

**LHS — Task Description**
> Complete `SumEvens(n int) int`. It should sum all even numbers from `1` to `n` (inclusive) using a `for` loop.
>
> **Expected behavior:**
> ```go
> SumEvens(10) // 2+4+6+8+10 = 30
> SumEvens(5)  // 2+4 = 6
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func SumEvens(n int) int {
	// TODO: sum all even numbers from 1 to n using a for loop
	return 0
}

func main() {
	fmt.Println(SumEvens(10)) // 30
	fmt.Println(SumEvens(5))  // 6
}
```

---

## Task 2.2b — `Factorial` (Intermediate)

**LHS — Task Description**
> Complete `Factorial(n int) int`. Use a `for` loop to return the product of every integer from `1` to `n`. By convention, `Factorial(0)` returns `1` (the empty product).
>
> **Expected behavior:**
> ```go
> Factorial(5) // 1*2*3*4*5 = 120
> Factorial(3) // 6
> Factorial(0) // 1
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Factorial(n int) int {
	// TODO: return 1*2*...*n using a for loop (0! == 1)
	return 0
}

func main() {
	fmt.Println(Factorial(5)) // 120
	fmt.Println(Factorial(3)) // 6
	fmt.Println(Factorial(0)) // 1
}
```

---

## Task 2.2c — `IsPrime` (Advanced)

**LHS — Task Description**
> Complete `IsPrime(n int) bool`. Return `true` if `n` is prime (greater than 1 and divisible only by 1 and itself), and `false` otherwise. For efficiency, only test divisors from `2` up to `i*i <= n` using a `for` loop with an early `return`.
>
> **Expected behavior:**
> ```go
> IsPrime(2)  // true
> IsPrime(7)  // true
> IsPrime(8)  // false
> IsPrime(1)  // false
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func IsPrime(n int) bool {
	// TODO: return true if n is prime, false otherwise
	// test divisors up to i*i <= n, return early on a divisor found
	return false
}

func main() {
	fmt.Println(IsPrime(2)) // true
	fmt.Println(IsPrime(7)) // true
	fmt.Println(IsPrime(8)) // false
	fmt.Println(IsPrime(1)) // false
}
```

---

## Task 2.3a — `DayType` (Beginner)

**LHS — Task Description**
> Complete `DayType(day string) string` using a `switch` statement. Return `"Weekend"` for `"Sat"` or `"Sun"`, and `"Weekday"` for anything else.
>
> **Expected behavior:**
> ```go
> DayType("Sat") // "Weekend"
> DayType("Mon") // "Weekday"
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func DayType(day string) string {
	// TODO: use switch to return "Weekend" for Sat/Sun, "Weekday" otherwise
	return ""
}

func main() {
	fmt.Println(DayType("Sat"))
	fmt.Println(DayType("Mon"))
}
```

---

## Task 2.3b — `DaysInMonth` (Intermediate)

**LHS — Task Description**
> Complete `DaysInMonth(m int) int` using a `switch` statement with grouped cases. Return the number of days for the given month number (`1`=January … `12`=December):
> - February → `28`
> - April, June, September, November → `30`
> - all others → `31`
> - an invalid month → `0`
>
> **Expected behavior:**
> ```go
> DaysInMonth(2)  // 28
> DaysInMonth(4)  // 30
> DaysInMonth(1)  // 31
> DaysInMonth(13) // 0
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func DaysInMonth(m int) int {
	// TODO: use a switch with grouped cases to return the day count
	return 0
}

func main() {
	fmt.Println(DaysInMonth(2))  // 28
	fmt.Println(DaysInMonth(4))  // 30
	fmt.Println(DaysInMonth(1))  // 31
	fmt.Println(DaysInMonth(13)) // 0
}
```

---

## Task 2.3c — `Season` (Advanced)

**LHS — Task Description**
> Complete `Season(m int) string` using an **expression-less `switch`** with grouped multi-value cases. Given a month number (`1`=January … `12`=December), return:
> - `"Winter"` for `12`, `1`, `2`
> - `"Spring"` for `3`, `4`, `5`
> - `"Summer"` for `6`, `7`, `8`
> - `"Fall"` for `9`, `10`, `11`
> - `"invalid"` for anything else
>
> **Expected behavior:**
> ```go
> Season(12) // "Winter"
> Season(4)  // "Spring"
> Season(8)  // "Summer"
> Season(10) // "Fall"
> Season(0)  // "invalid"
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Season(m int) string {
	// TODO: use an expression-less switch with grouped cases (e.g. case 12, 1, 2:)
	return ""
}

func main() {
	fmt.Println(Season(12)) // Winter
	fmt.Println(Season(4))  // Spring
	fmt.Println(Season(8))  // Summer
	fmt.Println(Season(10)) // Fall
	fmt.Println(Season(0))  // invalid
}
```

---

## Quick Reference — What Each Task's Hidden Test Checks

| Task | Test checks |
|------|-------------|
| 2.1a | `Grade(95) == "A"`, `Grade(82) == "B"`, `Grade(71) == "C"`, `Grade(50) == "F"` + edge cases |
| 2.1b | `Sign(42) == "positive"`, `Sign(-3) == "negative"`, `Sign(0) == "zero"` |
| 2.1c | `ClassFare(...)` matches fare rules for 5–6 combinations incl. student/senior overlap |
| 2.2a | `SumEvens(10) == 30`, `SumEvens(5) == 6` |
| 2.2b | `Factorial(5) == 120`, `Factorial(0) == 1` |
| 2.2c | `IsPrime(2) == true`, `IsPrime(7) == true`, `IsPrime(8) == false`, `IsPrime(1) == false` |
| 2.3a | `DayType("Sat") == "Weekend"`, `DayType("Mon") == "Weekday"` |
| 2.3b | `DaysInMonth(2) == 28`, `DaysInMonth(4) == 30`, `DaysInMonth(13) == 0` |
| 2.3c | `Season(12) == "Winter"`, `Season(4) == "Spring"`, `Season(8) == "Summer"`, `Season(10) == "Fall"`, `Season(0) == "invalid"` |
