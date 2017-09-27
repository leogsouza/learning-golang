package main

import "fmt"

func main() {

	var num1 int = 50
	var num2 int = 60

	if num1 != num2 && num1 <= num2 { // && Called Logical AND operator
		fmt.Println("True && operator")
	}

	if num1 != num2 || num1 <= num2 { // || Called Logical AND operator
		fmt.Println("True || operator")
	}

	if !(num1 == num2) {
		fmt.Println("True ! operator")
	}
}
