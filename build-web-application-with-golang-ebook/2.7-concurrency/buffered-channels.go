package main

import "fmt"

func main() {

	c := make(chan int, 3) // change 2 to 1 will have runtime error, but 3 is fine
	c <- 1
	c <- 2
	c <- 10

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
