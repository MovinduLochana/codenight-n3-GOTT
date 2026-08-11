package main

import "fmt"

func SetFive(ptr *int) {
	// TODO: set the variable pointed to by ptr to 5
}

func main() {
	x := 3
	SetFive(&x)
	fmt.Println(x) // 5
}