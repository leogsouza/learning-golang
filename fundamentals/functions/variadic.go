package main

import "fmt"

func add(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}

func main() {
	fmt.Println(add(1, 2, 3, 4))
}
