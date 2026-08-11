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