// Golang program to check whether a number is palindrome or not
package main

import "fmt"

func main() {
	var number, remainder, temp, reverse int

	fmt.Print("Enter any positive integer: ")
	fmt.Scan(&number)

	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		fmt.Println(reverse, remainder, number)
		number /= 10

		if number == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a Palindrome", temp)
	} else {
		fmt.Printf("%d is not a Palindrome", temp)
	}
}
