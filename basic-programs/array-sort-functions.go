package main

import (
	"fmt"
	"sort"
)

func main() {

	num := []int{60, 30, 70, 20, 35}
	fmt.Println(num)

	// Check if array of integer is sorted
	if sort.IntsAreSorted(num) == false {
		sort.Ints(num) // Sort integer Array
	}
	fmt.Println(num)
	fmt.Println(sort.SearchInts(num, 35))

	text := []string{"US", "UK", "Brazil", "Australia", "Canadá", "Austrália", "França", "Franca"}
	fmt.Println(text)

	if sort.StringsAreSorted(text) == false {
		sort.Strings(text)
	}
	fmt.Println(text)
	fmt.Println(sort.SearchStrings(text, "Austrália"))

	val := []float64{2.5, 3.1, 4.6, 2.7, 2.2, 3.8}
	fmt.Println(val)

	if sort.Float64sAreSorted(val) == false {
		sort.Float64s(val)
	}
	fmt.Println(val)
	fmt.Println(sort.SearchFloat64s(val, 2.2))
}
