package main

import "fmt"

func Receipt(item string, qty int, price float64) string {
	// TODO: use fmt.Sprintf to build and return the formatted receipt string
	return ""
}

func main() {
	fmt.Println(Receipt("Coffee", 2, 4.50)) // should print Item: Coffee, Qty: 2, Price: $4.50
}
