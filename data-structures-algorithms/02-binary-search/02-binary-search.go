package main

import "fmt"

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 34, 44, 55, 76, 81, 83, 85, 90, 92, 94, 98, 100, 110, 150, 170, 200, 203, 215, 225, 231, 244, 250}
	fmt.Println(newBinarySearch(250, items))
}
