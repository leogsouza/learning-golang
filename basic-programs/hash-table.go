// Example of Hash Table

package main

import "fmt"

func main() {

	var country map[int]string
	country = make(map[int]string)
	country[1] = "India"
	country[2] = "China"
	country[3] = "Brazil"
	country[4] = "France"
	country[5] = "Argentina"
	country[6] = "Cameroon"

	for i, j := range country {
		fmt.Printf("Key: %d Value: %s\n", i, j)
	}

}
