### Syntax
```go
go someFunc()      // run someFunc concurrently, don't block main
var wg sync.WaitGroup
wg.Add(1)          // register one goroutine we'll wait for
go func() { defer wg.Done(); /* work */ }()
wg.Wait()          // block until every Add'd goroutine calls Done
```

A **goroutine** is a lightweight, independently-running function — Go's version of a thread, but far cheaper (you can run thousands of them). Start one by putting `go` before a function call.

```go
go sayHello() // runs concurrently, doesn't block main()
```

Because goroutines run concurrently, you need a way to wait for them to finish before your program moves on (or exits). `sync.WaitGroup` is the standard tool for that:

```go
var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println("worker", id)
    }(i)
}
wg.Wait() // blocks until all 3 goroutines call Done()
```

> **Gotcha:** If multiple goroutines write to the same shared variable at the same time, you get a **data race**. Guard shared state with a `sync.Mutex`, as shown in the tasks below.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("hello from a goroutine")
    }()
    wg.Wait()
}
```