// This Program Calculate The Standard Deviation

package main

import (
	"fmt"
	"math"
)

func main() {
	var num [10]float64
	var sum, mean, sd float64

	fmt.Println("****** Enter 10 elements ******")

	for i := 1; i <= 10; i++ {
		fmt.Printf("Enter %d element: ", i)
		fmt.Scan(&num[i-1])
		sum += num[i-1]
	}

	mean = sum / 10

	for j := 0; j < 10; j++ {
		sd += math.Pow(num[j]-mean, 2)
	}

	sd = math.Sqrt(sd / 10)

	fmt.Println("The Standard Deviation is: ", sd)
}
