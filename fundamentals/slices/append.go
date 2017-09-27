package main

import "fmt"

func main() {

	// Create a smaller slice
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("Slice A:", a)
	fmt.Printf("Length is %d Capacity %d\n", len(a), cap(a))

	// Create a bigger slice
	a = append(a, 30, 40, 50)
	fmt.Println("Slice A after appending data:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
}
