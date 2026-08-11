package main

import "fmt"

// TODO: define `type Money float64`

// TODO: define a value-receiver method Format() float64 on Money
// that returns float64(m)

// TODO: define AddCurrency(a, b Money) Money

func main() {
	m := Money(12.5)
	fmt.Println(m.Format())                    // 12.5
	fmt.Println(AddCurrency(Money(1.25), Money(2.50))) // 3.75
}