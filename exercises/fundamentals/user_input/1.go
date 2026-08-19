package main

import "fmt"

func Greet() string {
	// TODO: read the user's name with fmt.Scanln, then return "Hello, <name>!"
	var name string
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanln(&name)

	if err != nil {
		return "Hello, Gopher!"
	}
	return "Hello, " + name + "!fgfhgfh"
}

func main() {
	fmt.Println(Greet()) // reads name, prints Hello, <name>!
}
