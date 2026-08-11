package main

import "fmt"

func StackLog() (log []string) {
	// TODO: append "start" normally

	// TODO: defer appending "B", THEN defer appending "A".
	// LIFO means "A" (deferred last) runs first.
	// Expected final log: ["start", "A", "B"]

	return
}

func main() {
	fmt.Println(StackLog())
}