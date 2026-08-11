Go supports the same three families of operators you'd expect from most languages:

**Arithmetic:** `+` `-` `*` `/` `%` (modulo — remainder after division)

```go
sum := 5 + 3       // 8
remainder := 7 % 2 // 1
```

**Comparison:** `==` `!=` `<` `>` `<=` `>=` — always return a `bool`

```go
isEqual := (5 == 5) // true
```

**Logical:** `&&` (and), `||` (or), `!` (not)

```go
canVote := age >= 18 && hasID
```

> **Gotcha:** Integer division **truncates** — it drops the decimal part entirely, it doesn't round.
> ```go
> result := 7 / 2 // 3, not 3.5
> ```
> We'll fix this properly in the next lesson on type conversion.

```go
package main

import "fmt"

func main() {
    n := 7
    isEven := n%2 == 0
    fmt.Println(isEven) // false
}
```