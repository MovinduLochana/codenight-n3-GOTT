package main

import "fmt"

// TODO: define interface Pricer with method Price() float64

type Product struct {
	Name      string
	BasePrice float64
}

func (p Product) Price() float64 {
	// TODO: return p.BasePrice
	return 0
}

func TotalPrice(items []Pricer) float64 {
	// TODO: sum every item's Price()
	return 0
}

func main() {
	items := []Pricer{
		Product{Name: "Coffee", BasePrice: 4.50},
		Product{Name: "Cake", BasePrice: 6.00},
	}
	fmt.Println(TotalPrice(items))
}