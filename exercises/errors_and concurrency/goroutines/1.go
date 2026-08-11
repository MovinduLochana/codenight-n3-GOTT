package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ids := []int{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		// TODO: launch `go func(id int) { ... }(i)` here.
		// Inside: defer wg.Done(), lock mu, append id to ids, unlock mu.
	}

	wg.Wait()
	sort.Ints(ids)   // sort just for predictable printing
	fmt.Println(ids) // [0 1 2]
}
