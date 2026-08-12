package main

import "fmt"

type Stack struct {
	elements []string
}

// TODO: Implement Push(val string) with a pointer receiver.

// TODO: Implement Pop() (string, bool) with a pointer receiver.
// It should return the popped string and true, or "", false if the stack is empty.

func main() {
	var s Stack

	fmt.Println("Beginning transaction...")
	s.Push("START")

	// TODO: Register a defer statement that will Pop and print all elements from the stack
	// until it is empty, demonstrating LIFO order.
	// Hint: The deferred block should print "Rolling back or auditing transaction steps:"
	// followed by each popped element prefixed with "- ".

	s.Push("WRITE_TEMP")
	s.Push("COMMIT")

	fmt.Println("Transaction steps registered.")
}
