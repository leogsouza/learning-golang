package main

import "fmt"

func main() {
	var action int
	fmt.Println("Enter 1 for Student and 2 for Professional")
	fmt.Scanln(&action)
	/* Use of Switch Case in Golang */
	switch action {
	case 1:
		fmt.Printf("I am a Student")
	case 2:
		fmt.Printf("I am a Professional")
	default:
		panic(fmt.Sprintf("I am a %d", action))
	}

	fmt.Println("")
	fmt.Println("Enter 1 for Student and 2 for UK")
	fmt.Scanln(&action)

	switch {
	case 1 == action:
		fmt.Printf("US")
	case 2 == action:
		fmt.Printf("UK")
	default:
		panic(fmt.Sprintf("I am a %d", action))
	}
}
