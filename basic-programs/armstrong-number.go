// Golang Program to check Armstrong Number

package main

import "fmt"

func main() {

	var number, tempNumber, remainder, result int

	fmt.Print("Enter any three digit number: ")
	fmt.Scan(&number)

	/* A positive integer is called an Armstrong number  of ofrder n if the sum of cubes of each digits is equal to the number
	   itself. 153 = 1*1*1 + 5*5*5 + 3*3*3 // 153 is an Armstrong number.
	*/

	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder
		tempNumber /= 10

		if tempNumber == 0 {
			break
		}
	}

	if result == number {
		fmt.Printf("%d is an Armstrong number.", number)
	} else {
		fmt.Printf("%d is not an Armstrong number.", number)
	}

}
