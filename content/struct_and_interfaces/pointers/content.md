### Syntax
```go
p := &x   // "address of" — p is a *T pointing at x
*p        // "dereference" — read the value p points to
*p = 10   // write THROUGH the pointer, mutating the original x
```

A **pointer** holds the *memory address* of a variable instead of a copy of its value. You create one with `&` ("address of") and read/write through it with `*` ("dereference").

```go
x := 42
p := &x    // p is a *int — a pointer to x
fmt.Println(p)  // 0xc0000b4008 (some address)
fmt.Println(*p) // 42 — dereference
```

The power of a pointer is that writing through it modifies the *original* variable:

```go
x := 7
p := &x
*p = 10 // writes through the pointer
fmt.Println(x) // 10 — x changed!
```

> **Key point:** A function that receives a pointer can mutate the caller's variable. A function that receives a plain value only gets a copy. This distinction is what makes **pointer receivers** in Go methods work — we'll build on it in Lesson 5.3.

```go
package main

import "fmt"

func main() {
    x := 7
    p := &x
    *p = 10
    fmt.Println(x) // 10
}
```