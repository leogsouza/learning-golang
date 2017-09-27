package main

import "fmt"

func main() {
	test1 := []int{5, 10, 15, 20, 25}
	fmt.Println(test1, len(test1), cap(test1))

	test2 := []int{}
	fmt.Println(test2, len(test2), cap(test2))

	test3 := []int{0: 9, 2: 18, 5: 27}
	fmt.Println(test3, len(test3), cap(test3))

	test4 := []int{7: 0}
	fmt.Println(test4, len(test4), cap(test4))
}
