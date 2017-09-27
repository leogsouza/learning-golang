package main

import "fmt"

func main() {

	var num []int
	var city []string
	var series []float64

	fmt.Println(num, len(num), cap(num))
	if num == nil {
		fmt.Println("Num is Nil")
	}

	fmt.Println(city, len(city), cap(city))
	if city == nil {
		fmt.Println("City is Nil")
	}

	fmt.Println(series, len(series), cap(series))
	if series == nil {
		fmt.Println("Series is Nil")
	}
}
