# Chapter 3 — Collections: Arrays, Slices, Maps
## Content Slides (Concept Screens)

Each slide below is the full-screen markdown shown to the student *before* the corresponding hands-on task screen. Each lesson follows the same structure: **Syntax → Explanation → Example → Callout → Full Runnable Snippet.**

---

## Lesson 3.1 — Arrays vs. Slices

### Syntax
```go
var arr [3]int = [3]int{1, 2, 3} // fixed-size array
nums := []int{1, 2, 3}           // slice — resizable
nums = append(nums, 4)           // grows the slice
len(s)                           // number of elements currently in s
cap(s)                           // capacity before a reallocation is needed
```

An **array** has a fixed size baked into its type — `[5]int` and `[10]int` are different types entirely. Arrays are rare in everyday Go code.

A **slice** is a flexible, resizable view over an underlying array, and it's what you'll use almost all the time.

```go
var arr [3]int = [3]int{1, 2, 3} // fixed size, always 3

nums := []int{1, 2, 3}           // slice — can grow
nums = append(nums, 4)           // now has 4 elements
```

`len(s)` returns the number of elements currently in a slice; `cap(s)` returns how many it can hold before Go needs to allocate a new backing array.

> **Key point:** When in doubt, use a slice, not an array. You'll almost never declare a fixed-size array in application code.

```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3}
    nums = append(nums, 4)
    fmt.Println(nums, len(nums)) // [1 2 3 4] 4
}
```

---

## Lesson 3.2 — Slicing `[low:high]` and `append`

### Syntax
```go
s[low:high]  // sub-slice from index low up to (not including) high
copy(dst, src) // copies elements into an independent slice
```

You can carve out a sub-slice using `slice[low:high]`:

```go
nums := []int{10, 20, 30, 40, 50}
middle := nums[1:3] // [20 30]
```

> **Gotcha:** A sub-slice shares the **same underlying array** as the original. Modifying one can silently modify the other — a classic source of bugs.
> ```go
> a := []int{1, 2, 3}
> b := a[0:2]
> b[0] = 99
> fmt.Println(a) // [99 2 3] — a changed too!
> ```
> When you need an independent copy, use `copy()`:
> ```go
> b := make([]int, len(a))
> copy(b, a)
> ```

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30, 40, 50}
    middle := nums[1:3]
    fmt.Println(middle) // [20 30]
}
```

---

## Lesson 3.3 — Maps

### Syntax
```go
m := make(map[string]int) // create an empty, writable map
m["key"] = 42             // set
v := m["key"]             // get
delete(m, "key")          // remove a key
v, ok := m["missing"]     // comma-ok: ok reports whether the key exists
```

A **map** stores key-value pairs, similar to a dictionary or hash table in other languages.

```go
ages := make(map[string]int)
ages["Alice"] = 30
ages["Bob"] = 25

fmt.Println(ages["Alice"]) // 30
```

To check whether a key actually exists (as opposed to just being absent and returning the zero value), use the **"comma ok" idiom**:

```go
value, ok := ages["Charlie"]
// value == 0, ok == false — "Charlie" isn't in the map
```

> **Gotcha:** A `nil` map (declared with `var m map[string]int` but never `make`'d) can be *read* safely, but writing to it causes a runtime panic. Always `make()` a map before writing to it.

```go
package main

import "fmt"

func main() {
    ages := make(map[string]int)
    ages["Alice"] = 30

    value, ok := ages["Bob"]
    fmt.Println(value, ok) // 0 false
}
```

---

## Lesson 3.4 — `range` over Slices & Maps

### Syntax
```go
for i, v := range nums { ... } // index + value
for _, v := range nums { ... } // value only (index discarded)
for k, v := range m { ... }    // map key + value
```

`range` is Go's way of iterating over slices, arrays, strings, and maps without manually managing an index.

```go
nums := []int{10, 20, 30}
for i, v := range nums {
    fmt.Println(i, v) // 0 10 / 1 20 / 2 30
}
```

If you don't need the index, discard it with `_`:

```go
for _, v := range nums {
    fmt.Println(v)
}
```

> **Gotcha:** Map iteration order is **randomized** by Go on purpose — every run may visit keys in a different order. Never rely on map order; sort the keys first if you need a consistent sequence.

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30}
    total := 0
    for _, v := range nums {
        total += v
    }
    fmt.Println(total) // 60
}
```
