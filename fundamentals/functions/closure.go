package main

import "fmt"

func main() {
	cube := func(x int) int {
		return x * x * x
	}

	fmt.Println(cube(5))
}
