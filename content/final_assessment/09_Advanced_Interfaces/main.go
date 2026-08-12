package main

import "fmt"

// TODO: Define the Item interface with a Price() float64 method.

// TODO: Define the BaseItem struct and implement its Price() method.

// TODO: Define the TaxableItem struct (embedding BaseItem) and override Price() to add 15% tax.

// TODO: Implement CalculateTotal(items []Item) float64.

func main() {
	book := BaseItem{Name: "Go Book", BasePrice: 30.00}
	laptop := TaxableItem{BaseItem: BaseItem{Name: "Laptop", BasePrice: 1000.00}}

	cart := []Item{book, laptop}
	fmt.Printf("Total Cart Price: $%.2f\n", CalculateTotal(cart))
}
