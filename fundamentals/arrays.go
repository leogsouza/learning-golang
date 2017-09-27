package main

import "fmt"

func main() {

	var x [5]int // Array Declaration

	x[0] = 10 // Assign the values to especific Index
	x[4] = 20 // Assign Value to array index in any Order
	x[1] = 30
	x[3] = 40
	x[2] = 50

	fmt.Println("Values of Array X: ", x)

	// Array Declaration and Initialization to specific Index
	y := [5]int{0: 100, 1: 200, 3: 500}

	fmt.Println("Values of Array Y: ", y)

	// Arrau Declaration and Initialization
	Country := [5]string{"US", "UK", "Australia", "Russia", "Brazil"}
	fmt.Println("Values of Array Country:", Country)

	// Array Declaration without length and Initialization
	Transport := [...]string{"Train", "Bus", "Plane", "Car", "Bike"}
	fmt.Println("Values of Array Transport", Transport)

}
