# Chapter 3 — Collections: Arrays, Slices, Maps
## Task Slides (Practice Screens)

Each task below is the split-screen practice slide shown right after the matching content slide.
**LHS** = task description shown to the student. **RHS** = starter code pre-loaded in the editor.
`// TODO` marks exactly where the student needs to write code.

---

## Task 3.1a — `AppendThree` (Beginner)

**LHS — Task Description**
> Complete `AppendThree(s []int) []int`. It should append the values `1`, `2`, and `3` (in that order) to the given slice and return the result.
>
> **Expected behavior:**
> ```go
> AppendThree([]int{})  // [1 2 3]
> AppendThree([]int{9}) // [9 1 2 3]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func AppendThree(s []int) []int {
	// TODO: append 1, 2, 3 to s and return it
	return s
}

func main() {
	fmt.Println(AppendThree([]int{}))
	fmt.Println(AppendThree([]int{9}))
}
```

---

## Task 3.1b — `Range` (Intermediate)

**LHS — Task Description**
> Complete `Range(n int) []int`. It should return a slice of integers from `0` up to (not including) `n`. Build the slice with `make([]int, n)` and fill it using a `for` loop.
>
> **Expected behavior:**
> ```go
> Range(5) // [0 1 2 3 4]
> Range(0) // []
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Range(n int) []int {
	// TODO: build a slice with make([]int, n), fill 0..n-1 with a for loop
	return nil
}

func main() {
	fmt.Println(Range(5)) // [0 1 2 3 4]
	fmt.Println(Range(0)) // []
}
```

---

## Task 3.1c — `Tail` (Advanced)

**LHS — Task Description**
> Complete `Tail(s []int, n int) []int`. It should return a **new independent copy** of the last `n` elements of `s` (without leaking extra capacity into the result). Use `make` + `copy`. If `n >= len(s)`, return a copy of the whole slice.
>
> **Expected behavior:**
> ```go
> Tail([]int{1, 2, 3, 4, 5}, 2) // [4 5]
> Tail([]int{1, 2}, 5)          // [1 2]
> Tail([]int{1, 2, 3}, 0)       // []
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Tail(s []int, n int) []int {
	// TODO: return an independent copy of the last n elements
	// use make + copy; if n >= len(s) copy the whole slice
	return nil
}

func main() {
	fmt.Println(Tail([]int{1, 2, 3, 4, 5}, 2)) // [4 5]
	fmt.Println(Tail([]int{1, 2}, 5))          // [1 2]
	fmt.Println(Tail([]int{1, 2, 3}, 0))       // []
}
```

---

## Task 3.2a — `Cut` (Beginner)

**LHS — Task Description**
> Complete `Cut(s []int, lo, hi int) []int`. It should return the sub-slice from index `lo` up to (not including) `hi`, using slice notation `s[lo:hi]`.
>
> **Expected behavior:**
> ```go
> Cut([]int{10, 20, 30, 40, 50}, 1, 3) // [20 30]
> Cut([]int{5, 6, 7}, 0, 2)            // [5 6]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Cut(s []int, lo, hi int) []int {
	// TODO: return the sub-slice s[lo:hi]
	return nil
}

func main() {
	fmt.Println(Cut([]int{10, 20, 30, 40, 50}, 1, 3)) // [20 30]
	fmt.Println(Cut([]int{5, 6, 7}, 0, 2))            // [5 6]
}
```

---

## Task 3.2b — `RemoveAt` (Intermediate)

**LHS — Task Description**
> Complete `RemoveAt(s []int, i int) []int`. It should return a new slice with the element at index `i` removed, keeping the remaining order intact. The idiomatic one-liner is `append(s[:i], s[i+1:]...)`.
>
> > **Gotcha:** this mutates `s`'s backing array — the caller's original slice is changed too. That's expected here; we're demonstrating shared-array behavior, not defending against it.
>
> **Expected behavior:**
> ```go
> RemoveAt([]int{10, 20, 30}, 1) // [10 30]
> RemoveAt([]int{5, 6, 7, 8}, 0) // [6 7 8]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func RemoveAt(s []int, i int) []int {
	// TODO: return a new slice with the element at index i removed
	return s
}

func main() {
	fmt.Println(RemoveAt([]int{10, 20, 30}, 1)) // [10 30]
	fmt.Println(RemoveAt([]int{5, 6, 7, 8}, 0)) // [6 7 8]
}
```

---

## Task 3.2c — `InsertAt` (Advanced)

**LHS — Task Description**
> Complete `InsertAt(s []int, i, v int) []int`. It should return a new slice identical to `s` but with `v` inserted at index `i`, preserving order. Idiomatic approach: lengthen `s` by one with `append`, shift the tail right with `copy`, then set index `i` to `v`. Assume `i` is a valid index and `s` is non-empty.
>
> **Expected behavior:**
> ```go
> InsertAt([]int{10, 20, 30}, 1, 99) // [10 99 20 30]
> InsertAt([]int{1, 2}, 0, 0)        // [0 1 2]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func InsertAt(s []int, i, v int) []int {
	// TODO: return a new slice with v inserted at index i, order preserved
	// hint: append a zero value to grow s, then re-slice and shift
	return nil
}

func main() {
	fmt.Println(InsertAt([]int{10, 20, 30}, 1, 99)) // [10 99 20 30]
	fmt.Println(InsertAt([]int{1, 2}, 0, 0))        // [0 1 2]
}
```

---

## Task 3.3a — `WordCount` (Beginner)

**LHS — Task Description**
> Complete `WordCount(words []string) map[string]int`. It should return a map where each key is a word and each value is how many times that word appeared in the input slice.
>
> **Expected behavior (map order not guaranteed — hidden test compares map equality):**
> ```go
> WordCount([]string{"go", "is", "fun", "go", "is", "go"})
> // map[fun:1 go:3 is:2]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func WordCount(words []string) map[string]int {
	// TODO: build and return a map counting occurrences of each word
	return nil
}

func main() {
	fmt.Println(WordCount([]string{"go", "is", "fun", "go", "is", "go"}))
}
```

---

## Task 3.3b — `MergeCounts` (Intermediate)

**LHS — Task Description**
> Complete `MergeCounts(a, b map[string]int) map[string]int`. It should return a new map that combines both input maps by summing the count for every key present in either. Use the **comma-ok idiom** to read from `b` and determine whether a key exists.
>
> **Expected behavior (map order not guaranteed — hidden test compares map equality):**
> ```go
> MergeCounts(
>     map[string]int{"go": 2, "is": 1},
>     map[string]int{"go": 1, "fun": 3},
> ) // map[fun:3 go:3 is:1]
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func MergeCounts(a, b map[string]int) map[string]int {
	// TODO: combine a and b into a new map, summing counts
	return nil
}

func main() {
	fmt.Println(MergeCounts(
		map[string]int{"go": 2, "is": 1},
		map[string]int{"go": 1, "fun": 3},
	))
}
```

---

## Task 3.3c — `Invert` (Advanced)

**LHS — Task Description**
> Complete `Invert(m map[string]int) map[int][]string`. It should return a map that swaps keys and values: each original value becomes a key, and its slice holds all the original keys that mapped to it. Build the `[]string` with `make`/`append`.
>
> **Expected behavior (order of each value slice not guaranteed — hidden test deep-compares):**
> ```go
> Invert(map[string]int{"a": 1, "b": 1, "c": 2})
> // map[1:[a b] 2:[c]]  (order of a/b varies)
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Invert(m map[string]int) map[int][]string {
	// TODO: return a map whose keys are m's values and whose values
	// are slices of m's keys that mapped to that value
	return nil
}

func main() {
	fmt.Println(Invert(map[string]int{"a": 1, "b": 1, "c": 2}))
}
```

---

## Task 3.4a — `Sum` (Beginner)

**LHS — Task Description**
> Complete `Sum(nums []int) int` using a `range` loop (not indexing with `[]`). It should return the sum of all elements.
>
> **Expected behavior:**
> ```go
> Sum([]int{1, 2, 3, 4}) // 10
> Sum([]int{})           // 0
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Sum(nums []int) int {
	// TODO: use range to sum all elements in nums
	return 0
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4}))
	fmt.Println(Sum([]int{}))
}
```

---

## Task 3.4b — `Contains` (Intermediate)

**LHS — Task Description**
> Complete `Contains(nums []int, target int) bool`. It should return `true` if `target` appears anywhere in `nums`. Use a `range` loop that discards the index with `_` and returns early on a match.
>
> **Expected behavior:**
> ```go
> Contains([]int{1, 2, 3}, 2) // true
> Contains([]int{1, 2, 3}, 9) // false
> Contains([]int{}, 0)        // false
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func Contains(nums []int, target int) bool {
	// TODO: use range (ignore index) to find target, return early
	return false
}

func main() {
	fmt.Println(Contains([]int{1, 2, 3}, 2)) // true
	fmt.Println(Contains([]int{1, 2, 3}, 9)) // false
	fmt.Println(Contains([]int{}, 0))        // false
}
```

---

## Task 3.4c — `IndexAll` (Advanced)

**LHS — Task Description**
> Complete `IndexAll(nums []int, target int) []int`. It should return a slice holding the index of every position where `target` appears. Build the result with `append` inside a `range` loop that uses the index (`i`).
>
> **Expected behavior:**
> ```go
> IndexAll([]int{1, 2, 1, 4, 1}, 1) // [0 2 4]
> IndexAll([]int{1, 2, 3}, 9)       // []
> ```

**RHS — Starter Code**
```go
package main

import "fmt"

func IndexAll(nums []int, target int) []int {
	// TODO: use range with index to collect every position of target
	return nil
}

func main() {
	fmt.Println(IndexAll([]int{1, 2, 1, 4, 1}, 1)) // [0 2 4]
	fmt.Println(IndexAll([]int{1, 2, 3}, 9))       // []
}
```

---

## Quick Reference — What Each Task's Hidden Test Checks

| Task | Test checks |
|------|-------------|
| 3.1a | `AppendThree([]int{}) == [1 2 3]`, `AppendThree([]int{9}) == [9 1 2 3]` |
| 3.1b | `Range(5) == [0 1 2 3 4]`, `Range(0) == []` |
| 3.1c | `Tail(...)` returns an independent copy with exact `len` and no extra `cap` |
| 3.2a | `Cut(...)` returns correct sub-slices for 2–3 inputs |
| 3.2b | `RemoveAt(...)` == `[10 30]` / `[6 7 8]` |
| 3.2c | `InsertAt(...)` == `[10 99 20 30]` / `[0 1 2]` |
| 3.3a | map equality with `map[fun:1 go:3 is:2]` |
| 3.3b | map equality with `map[fun:3 go:3 is:1]` |
| 3.3c | deep map equality (value slices compared as sets) |
| 3.4a | `Sum([]int{1,2,3,4}) == 10`, `Sum([]int{}) == 0` |
| 3.4b | `Contains(...)` true/false for 3 cases |
| 3.4c | `IndexAll(...) == [0 2 4]`, empty for no match |