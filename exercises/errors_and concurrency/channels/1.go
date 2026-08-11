package main

import "fmt"

func Producer(ch chan<- int, n int) {
	// TODO: send 1 through n into ch, then close(ch)
}

func main() {
	ch := make(chan int)
	go Producer(ch, 3)

	for v := range ch {
		fmt.Println(v)
	}
}
