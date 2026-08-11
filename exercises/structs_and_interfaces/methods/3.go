package main

import "fmt"

// TODO: define `type Stack struct { Items []int }`

// TODO: define Push(v int) on *Stack — append v to Items

// TODO: define Pop() int on *Stack — remove and return the last item

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop()) // 3
	fmt.Println(s.Pop()) // 2
	fmt.Println(s.Items) // [1]
}