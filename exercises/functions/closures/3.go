package main

import "fmt"

func MakeBank(balance float64) func(amount float64) float64 {
	// TODO: return a closure capturing `balance` that:
	//   - deposits positive amounts and returns the new balance
	//   - ignores zero/negative amounts (returns current balance unchanged)
	return nil
}

func main() {
	bank := MakeBank(100)
	fmt.Println(bank(50))  // 150
	fmt.Println(bank(-20)) // 150
	fmt.Println(bank(25))  // 175
}