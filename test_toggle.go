package main

import (
	"fmt"
)

func main() {

	var input []byte
	fmt.Scan(&input)
	for i := 0; i < len(input); i++ {
		if input[i] >= 'A' && input[i] <= 'Z' {
			input[i] += 32
		} else if input[i] >= 'a' && input[i] <= 'z' {
			input[i] -= 32
		}
	}

	fmt.Println(string(input))
}
