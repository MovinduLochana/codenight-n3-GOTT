package main

import "fmt"

// TODO: define Product struct with Name string and Price float64
// TODO: give Product a Label() string method returning
//   fmt.Sprintf("%s: $%.2f", p.Name, p.Price)

func main() {
	fmt.Println(Product{Name: "Coffee", Price: 4.50}.Label())
	fmt.Println(Product{Name: "Tea", Price: 3.00}.Label())
}