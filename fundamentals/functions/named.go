package main

import "fmt"

// function name add
func add(a int, b int) int { // int type of function
	var c int
	c = a + b
	return c // return type
}

func main() {
	x := 10
	y := 20
	z := add(x, y)

	fmt.Println(z)
}
