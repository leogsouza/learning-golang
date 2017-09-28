// This program print the largest element of an Array

package main

import "fmt"

func main() {
	var num [100]float64
	var temp int

	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&temp)

	for i := 0; i < temp; i++ {
		fmt.Print("Enter the number: ")
		fmt.Scanln(&num[i])
	}

	for j := 1; j < temp; j++ {
		if num[0] < num[j] {
			num[0] = num[j]
		}
	}

	fmt.Print("The largest number is: ", num[0])
}
