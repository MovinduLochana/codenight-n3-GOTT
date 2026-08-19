Everything so far has been one-way: your program prints output and the user just watches. But real programs are **interactive** — they read what the user types and respond to it. Go's `fmt` package gives you a matching set of **input** functions to go along with the print functions you already know.

```go
var name string
fmt.Println("What's your name?")
fmt.Scanln(&name)          // reads one line of input into name
fmt.Println("Hello,", name)
```

Two you'll use most often:

| Function | Reads | Stops at |
|----------|-------|----------|
| `fmt.Scan(&a, &b)` | space-separated values | whitespace |
| `fmt.Scanln(&a, &b)` | space-separated values | a newline |

Both take **pointers** (`&name`, `&age`) — that's how the function writes the value back into your variable. That `&` "address-of" operator is the same one you use when a function needs to modify a variable.

```go
var age int
fmt.Print("Enter your age: ")
fmt.Scanln(&age)          // type must match: int expects an integer
fmt.Printf("Next year you'll be %d!\n", age+1)
```

> **Key point:** `Scan`/`Scanln` fill the variables you pass in — that's why they need `&`. Get the types wrong (e.g. typing `abc` into an `int`) and the scan stops early, leaving your variables at their zero value.

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("What's your name? ")
    fmt.Scanln(&name)
    fmt.Println("Hello,", name)
}
```