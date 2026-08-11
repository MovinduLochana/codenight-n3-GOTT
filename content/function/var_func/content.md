### Syntax
```go
func <name>(<params>, last ...T) { ... } // last param is a []T
sum(slice...)                            // spread a slice into a variadic call
```

A **variadic** function accepts a variable number of arguments of the same type — you've already used one every time you called `fmt.Println`.

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)       // works with any number of args
sum()              // even zero args
```

Inside the function, `nums` behaves exactly like a regular `[]int` slice. You can also "spread" an existing slice into a variadic call using `...`:

```go
values := []int{1, 2, 3}
sum(values...)
```

> **Key point:** A variadic parameter must be the **last** parameter in the function signature, and a function can only have one.

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3)) // 6
}
```