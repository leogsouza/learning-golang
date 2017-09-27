package main

import "fmt"

func cube() func() int {
	var x int
	return func() int {
		x++
		return x * x * x
	}
}

func main() {
	f := cube()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
