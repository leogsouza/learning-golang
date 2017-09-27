package main

import (
	"fmt"
)

func main() {

	// Create a smaller slice
	a := []int{5, 6, 7}
	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))
	// Create a bigger slice
	b := make([]int, 5, 10)
	copy(b, a) // copy function
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

	fmt.Println("Slice B after copying:", b)
	b[3] = 8
	b[4] = 9
	fmt.Println("Slice B after adding elements:", b)
}
