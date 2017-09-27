package main

import "fmt"

func main() {
    
    var country map[int]string // Declaration of Nil map
    country = make(map[int] string) // Initialization of map using make function
    country[1] = "India"
    country[2] = "China"
    country[3] = "Pakistan"
    country[4] = "Germany"
    country[5] = "Australia"
    country[6] = "Indonesia"
    
    // Run a loop to fetch map elements using range
    for i, j := range country {
        fmt.Printf("Key: %d Value: %s\n", i, j)
    }
}