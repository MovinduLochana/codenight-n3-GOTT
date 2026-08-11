package main

import "fmt"

type Pricer interface {
	Price() float64
}

type Product struct {
	Name      string
	BasePrice float64
}

func (p Product) Price() float64 {
	return p.BasePrice
}

// TODO: define Deal embedding Product, plus a Discount float64 field
// TODO: give Deal its OWN Price() float64 method returning
//   d.Product.BasePrice * (1 - d.Discount)

func TotalPrice(items []Pricer) float64 {
	// TODO: sum every item's Price()
	return 0
}

func main() {
	items := []Pricer{
		Product{Name: "Coffee", BasePrice: 4.00},
		Deal{Product: Product{Name: "Cake", BasePrice: 10.00}, Discount: 0.20},
	}
	fmt.Println(TotalPrice(items))
}