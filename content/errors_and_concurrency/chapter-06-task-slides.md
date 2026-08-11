# Chapter 6 — Errors & a Taste of Concurrency
## Task Slides (Practice Screens)

Each task below is the split-screen practice slide shown right after the matching content slide.
**LHS** = task description shown to the student. **RHS** = starter code pre-loaded in the editor.
`// TODO` marks exactly where the student needs to write code.

---

## Task 6.1a — `ParsePositive` (Beginner)

**LHS — Task Description**
> Complete `ParsePositive(s string) (int, error)`. Use `strconv.Atoi` to parse `s`. Return an error if parsing fails, or if the parsed number is negative. Otherwise return the number and `nil`.
>
> **Expected behavior:**
> ```go
> ParsePositive("42")  // 42, nil
> ParsePositive("-5")  // 0, error
> ParsePositive("abc") // 0, error
> ```

**RHS — Starter Code**
```go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func ParsePositive(s string) (int, error) {
	// TODO: parse s with strconv.Atoi.
	// Return the parse error unwrapped if it fails.
	// Return errors.New("number must be positive") if n < 0.
	// Otherwise return n, nil.
	return 0, nil
}

func main() {
	n, err := ParsePositive("42")
	fmt.Println(n, err)

	n, err = ParsePositive("-5")
	fmt.Println(n, err)

	n, err = ParsePositive("abc")
	fmt.Println(n, err)
}
```

---

## Task 6.1b — `ReadAge` (Intermediate)

**LHS — Task Description**
> Complete `ReadAge(s string) (int, error)`. Parse `s` with `strconv.Atoi`. On failure, **wrap** the error with `fmt.Errorf("invalid age: %w", err)` — the `%w` verb preserves the original error so `errors.Is(err, strconv.ErrSyntax)` works.
>
> **Expected behavior:**
> ```go
> ReadAge("25") // 25, nil
> ReadAge("xx") // 0, error wrapped with "invalid age: ..."
> errors.Is(err, strconv.ErrSyntax) // true
> ```

**RHS — Starter Code**
```go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func ReadAge(s string) (int, error) {
	// TODO: parse s with strconv.Atoi.
	// On failure, wrap the error: fmt.Errorf("invalid age: %w", err)
	return 0, nil
}

func main() {
	n, err := ReadAge("25")
	fmt.Println(n, err)

	_, err = ReadAge("xx")
	fmt.Println(err)                               // invalid age: strconv.Atoi: ...
	fmt.Println(errors.Is(err, strconv.ErrSyntax)) // true
}
```

---

## Task 6.1c — `ParsePositive` with a sentinel (Advanced)

**LHS — Task Description**
> Refactor with a **sentinel error**: declare `var ErrNotPositive = errors.New("number must be positive")` as a reusable, comparable value. In `ParsePositive`, return `0, ErrNotPositive` whenever `n <= 0`; return the parse error **unwrapped** on failure. Callers distinguish cases with `errors.Is`.
>
> **Expected behavior:**
> ```go
> ParsePositive("7")  // 7, nil
> errors.Is(err, ErrNotPositive)          // false
> ParsePositive("-5")                      // 0, ErrNotPositive
> errors.Is(ParsePositive("0"))            // true
> ParsePositive("abc") // strconv.Atoi error, NOT wrapped
> ```

**RHS — Starter Code**
```go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// TODO: declare `var ErrNotPositive = errors.New("number must be positive")`

func ParsePositive(s string) (int, error) {
	// TODO: parse s with strconv.Atoi (return parse error unwrapped).
	// If n <= 0, return 0, ErrNotPositive. Otherwise return n, nil.
	return 0, nil
}

func main() {
	n, err := ParsePositive("7")
	fmt.Println(n, err)
	fmt.Println(errors.Is(err, ErrNotPositive)) // false

	_, err = ParsePositive("-5")
	fmt.Println(errors.Is(err, ErrNotPositive)) // true

	_, err = ParsePositive("0")
	fmt.Println(errors.Is(err, ErrNotPositive)) // true

	_, err = ParsePositive("abc")
	fmt.Println(err)                            // strconv.Atoi error
	fmt.Println(errors.Is(err, ErrNotPositive)) // false
}
```

---

## Task 6.2a — Concurrent ID appends (Beginner)

**LHS — Task Description**
> Launch 3 goroutines (with IDs `0`, `1`, `2`), each appending its ID to the shared `ids` slice. The `sync.WaitGroup` and `sync.Mutex` are already wired up for you — you only need to write the `go func(id int) { ... }(i)` call that locks the mutex, appends `id`, and calls `wg.Done()`.
>
> **Expected behavior:** after `wg.Wait()`, `ids` contains `0`, `1`, and `2` in some order (goroutine order isn't guaranteed).

**RHS — Starter Code**
```go
package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ids := []int{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		// TODO: launch `go func(id int) { ... }(i)` here.
		// Inside: defer wg.Done(), lock mu, append id to ids, unlock mu.
	}

	wg.Wait()
	sort.Ints(ids)   // sort just for predictable printing
	fmt.Println(ids) // [0 1 2]
}
```

---

## Task 6.2b — `ConcurrentSum` (Intermediate)

**LHS — Task Description**
> Complete `ConcurrentSum(nums []int) int`. Launch a goroutine for **each** number; each goroutine adds its number to a shared `total` guarded by a mutex. A `sync.WaitGroup` waits for all of them. Return the final total. Order does not matter — only correctness under concurrency.
>
> **Expected behavior:**
> ```go
> ConcurrentSum([]int{1, 2, 3, 4}) // 10
> ConcurrentSum([]int{})           // 0
> ```

**RHS — Starter Code**
```go
package main

import (
	"fmt"
	"sync"
)

func ConcurrentSum(nums []int) int {
	// TODO: for each number in nums, launch a goroutine that adds it
	// to a shared `total` guarded by a sync.Mutex.
	// Use a sync.WaitGroup; return the final total.
	return 0
}

func main() {
	fmt.Println(ConcurrentSum([]int{1, 2, 3, 4})) // 10
	fmt.Println(ConcurrentSum([]int{}))           // 0
}
```

---

## Task 6.2c — `Squares` (Advanced)

**LHS — Task Description**
> Complete `Squares(nums []int) []int`. Every number should be squared **in its own goroutine** and the result stored into a shared result slice guarded by a mutex. Because goroutine completion order is unpredictable, **sort** the result before returning so the output is deterministic. Return a slice of the same length as `nums`, where every element is `nums[i] * nums[i]`.
>
> **Expected behavior (order guaranteed after sorting):**
> ```go
> Squares([]int{4, 1, 3}) // [1 9 16]
> Squares([]int{})        // []
> ```

**RHS — Starter Code**
```go
package main

import (
	"fmt"
	"sync"
)

func Squares(nums []int) []int {
	// TODO: launch one goroutine per number; each squares its value and
	// stores it into a shared result slice guarded by a sync.Mutex.
	// Wait for all goroutines, then SORT the result before returning.
	return nil
}

func main() {
	fmt.Println(Squares([]int{4, 1, 3})) // [1 9 16]
	fmt.Println(Squares([]int{}))        // []
}
```

---

## Task 6.3a — `Producer` (Beginner)

**LHS — Task Description**
> Complete `Producer(ch chan<- int, n int)`. It should send the numbers `1` through `n` (inclusive) into `ch`, in order, then close the channel. The consumer code in `main` is already written for you.
>
> **Expected behavior:**
> ```go
> // Producer(ch, 3) sends 1, 2, 3 then closes ch
> // consumer prints: 1 2 3
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Producer(ch chan<- int, n int) {
	// TODO: send 1 through n into ch, then close(ch)
}

func main() {
	ch := make(chan int)
	go Producer(ch, 3)

	for v := range ch {
		fmt.Println(v)
	}
}
```

---

## Task 6.3b — `SquaresPipeline` (Intermediate)

**LHS — Task Description**
> Complete `SquaresPipeline(k int) []int`. Build a two-goroutine **pipeline**: a producer sends `1..k` into a first channel, a transform goroutine reads it, squares each value, and sends the result into a second channel (then closes it). `main` drains the second channel with `for v := range ch` into the returned slice.
>
> **Expected behavior:**
> ```go
> SquaresPipeline(4) // [1 4 9 16]
> SquaresPipeline(0) // []
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func SquaresPipeline(k int) []int {
	// TODO: build a two-stage pipeline with channels:
	//   1. producer goroutine sends 1..k into `nums`
	//   2. transformer goroutine reads nums, squares, sends into `squares`
	//   3. main drains squares with for range, appending to the result
	return nil
}

func main() {
	fmt.Println(SquaresPipeline(4)) // [1 4 9 16]
	fmt.Println(SquaresPipeline(0)) // []
}
```

---

## Task 6.3c — `EvenSum` (Advanced)

**LHS — Task Description**
> Complete `EvenSum(n int) int`. Launch a goroutine that sends every even number from `1` to `n` (inclusive) into a channel, then **closes** it. Back in `main`, sum the values using `for v := range ch`. Return the total.
>
> **Expected behavior:**
> ```go
> EvenSum(10) // 2+4+6+8+10 = 30
> EvenSum(5)  // 2+4 = 6
> EvenSum(0)  // 0
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func EvenSum(n int) int {
	// TODO: launch a goroutine that sends every even number 1..n
	// into a channel, then closes it.
	// In main, sum the values with `for v := range ch`.
	return 0
}

func main() {
	fmt.Println(EvenSum(10)) // 30
	fmt.Println(EvenSum(5))  // 6
	fmt.Println(EvenSum(0))  // 0
}
```

---

## Task 6.4a — `Validate` (Beginner)

**LHS — Task Description**
> Complete `Validate(item string, qty int) error`. Return an error (using `errors.New` or `fmt.Errorf`) when `qty < 1` — an error **is** the return value here. Return `nil` when the quantity is valid. This is the "errors as values" pattern from Lesson 6.1.
>
> **Expected behavior:**
> ```go
> Validate("Coffee", 2) // nil
> Validate("Tea", 0)    // error
> Validate("Tea", -3)   // error
> ```

**RHS — Starter Code**
```go
package main

import (
	"errors"
	"fmt"
)

func Validate(item string, qty int) error {
	// TODO: return an error when qty < 1, otherwise nil
	return nil
}

func main() {
	fmt.Println(Validate("Coffee", 2))
	fmt.Println(Validate("Tea", 0))
	fmt.Println(Validate("Tea", -3))
}
```

---

## Task 6.4b — `CheckoutTotals` (Intermediate)

**LHS — Task Description**
> Complete `CheckoutTotals() int`. Launch a goroutine that sends the totals `10`, `20`, `30` (in order) into the `totals` channel and then **closes** it — the classic producer. Wait for it with a `sync.WaitGroup`. Back in `main`, sum the values with `for v := range totals` and return the total.
>
> **Note:** the channel is buffered (`make(chan int, 3)`) so the producer can finish without blocking — the unbuffered variant would deadlock the `wg.Wait()` above.
>
> **Expected behavior:**
> ```go
> CheckoutTotals() // 60
> ```

**RHS — Starter Code**
```go
package main

import (
	"fmt"
	"sync"
)

func CheckoutTotals() int {
	// Buffered so the producer can send all 3 values and close without
	// blocking — the unbuffered variant would deadlock the wg.Wait() above.
	totals := make(chan int, 3)

	var wg sync.WaitGroup
	wg.Add(1)
	// TODO: launch a goroutine (inside, defer wg.Done()) that sends
	// 10, 20, 30 into `totals`, then close(totals)

	wg.Wait()

	// TODO: sum the values with `for v := range totals` and return the total
	return 0
}

func main() {
	fmt.Println(CheckoutTotals()) // 60
}
```

---

## Task 6.4c — `Checkout` (Advanced)

**LHS — Task Description**
> Complete `Checkout(items []string, qty []int) ([]int, error)`. It should validate every `(item, qty)` pair with `Validate` (from the beginner task): if any quantity is invalid, return `nil, error` immediately. Otherwise compute each `qty[i] * 2` (a fake per-item price) **in its own goroutine**, send the results through a channel, and collect them into a slice with `for v := range ch`. A `WaitGroup` coordinates the workers. **Sort** the collected slice before returning so the output is deterministic despite goroutine ordering.
>
> **Expected behavior:**
> ```go
> Checkout([]string{"Coffee", "Tea"}, []int{2, 3}) // [4 6], nil
> Checkout([]string{"Coffee"}, []int{0})           // nil, error
> ```

**RHS — Starter Code**
```go
package main

import (
	"fmt"
	"sync"
)

func Validate(item string, qty int) error {
	if qty < 1 {
		return fmt.Errorf("invalid quantity %d for %s", qty, item)
	}
	return nil
}

func Checkout(items []string, qty []int) ([]int, error) {
	// TODO: validate every (item, qty) pair with Validate, return nil, err
	// if any is invalid.

	// TODO: launch one goroutine per item that computes qty[i] * 2 and
	// sends the result into a channel. Use a WaitGroup. Wait, then collect
	// all values with `for v := range ch`, SORT the result, and return it.
	return nil, nil
}

func main() {
	fmt.Println(Checkout([]string{"Coffee", "Tea"}, []int{2, 3}))
	fmt.Println(Checkout([]string{"Coffee"}, []int{0}))
}
```

---

## Quick Reference — What Each Task's Hidden Test Checks

| Task | Test checks |
|------|-------------|
| 6.1a | `ParsePositive("42")` → `42, nil`; `"-5"`/`"abc"` → error |
| 6.1b | `ReadAge("xx")` error wraps `strconv.ErrSyntax` (`errors.Is` true) |
| 6.1c | `-5`/`0` both `errors.Is(ErrNotPositive)`; `"abc"` returns unwrapped Atoi error |
| 6.2a | `ids` contains `{0, 1, 2}` after `wg.Wait()` (any order) |
| 6.2b | `ConcurrentSum([]int{1,2,3,4}) == 10`, `ConcurrentSum([]int{}) == 0` |
| 6.2c | `Squares([]int{4,1,3}) == [1 9 16]`, `Squares([]int{}) == []` |
| 6.3a | stdout is `1 2 3` |
| 6.3b | `SquaresPipeline(4) == [1 4 9 16]`, `SquaresPipeline(0) == []` |
| 6.3c | `EvenSum(10) == 30`, `EvenSum(5) == 6`, `EvenSum(0) == 0` |
| 6.4a | `Validate("Tea", 0)` / `("Tea", -3)` return errors; valid qty returns nil |
| 6.4b | `CheckoutTotals() == 60` |
| 6.4c | valid input → sorted `[4 6], nil`; invalid qty → `nil, error` |