// GO Program with Array Reverse

package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Integer Reverse Sort")
	num := []int{50, 90, 30, 10, 50}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)

	fmt.Println()

	fmt.Println("String Reverse Sort")
	text := []string{"Japan", "UK", "Canada", "US", "Germany", "India"}
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)
}
