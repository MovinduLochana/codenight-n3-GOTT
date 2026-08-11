package main

import "fmt"

const (
	_ = iota
	Red
	Green
	Blue
	Alpha
)

func BlueValue() int {
	// TODO: return the value of Blue
	return 0
}

func main() {
	fmt.Println(BlueValue())
}
