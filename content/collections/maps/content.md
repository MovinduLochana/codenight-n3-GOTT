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