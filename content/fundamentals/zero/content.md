What happens if you declare a variable but don't give it a value? Unlike some languages, Go doesn't leave it as garbage or `undefined`. Every type has a **zero value** — a sensible default it's automatically set to.

| Type | Zero value |
|------|-----------|
| `int`, `float64` | `0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| pointers, slices, maps, etc. | `nil` |

```go
var score int
var passed bool

fmt.Println(score, passed) // 0 false
```

> **Key point:** Zero values mean a Go variable is *always* in a valid, usable state the moment it's declared — you never have to worry about reading "uninitialized memory."

```go
package main

import "fmt"

func main() {
    var count int
    var message string
    var isReady bool

    fmt.Println(count, message, isReady) // 0 "" false
}
```