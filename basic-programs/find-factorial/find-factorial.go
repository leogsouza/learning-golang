package main

import "fmt"

var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
// Range 0 through 18446744073709551615
var i int = 1
var n int

func factorial(n int) uint64 {

	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i) // mismatched types int64 and int
		}
	}

	return factVal
}

func main() {
	fmt.Print("Enter a positive number between 0 - 50: ")
	fmt.Scan(&n)
	fmt.Print("Factorial is: ", factorial(n))
}
