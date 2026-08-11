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