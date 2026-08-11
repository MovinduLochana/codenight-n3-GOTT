package main

import "fmt"

func IsPrime(n int) bool {
	// TODO: return true if n is prime, false otherwise
	// test divisors up to i*i <= n, return early on a divisor found
	return false
}

func main() {
	fmt.Println(IsPrime(2)) // true
	fmt.Println(IsPrime(7)) // true
	fmt.Println(IsPrime(8)) // false
	fmt.Println(IsPrime(1)) // false
}