package main

import "fmt"

// TODO: define the Product struct (Name string, Price float64)
// TODO: define the Cart struct with field Items []Product

func AddItem(c Cart, p Product) Cart {
	// TODO: return a new Cart with p appended to c.Items
	return c
}

func main() {
	c := Cart{}
	c = AddItem(c, Product{Name: "Coffee", Price: 4.50})
	c = AddItem(c, Product{Name: "Tea", Price: 3.00})
	fmt.Println(len(c.Items), c.Items[0].Name)
}