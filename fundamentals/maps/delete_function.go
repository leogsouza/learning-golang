package main

import "fmt"

func main() {

	books := map[string]string{
		"Khushwant Singh": "A History of Sikhs",
		"Harisen":         "Allahabad Prasasti",
		"Akhil Sharma":    "An Obedient Father",
		"Premendra Mitra": "Agamikal",
		"Vikram Seth":     "An Equal Music", // This is required on Go
	}

	for writers := range books {
		fmt.Println("Book:", books[writers], "-Writer-", writers)
	}

	// Delete an item from map
	delete(books, "Harisen")

	// Print map again
	fmt.Println("")
	for writers := range books {
		fmt.Println("Book:", books[writers], "-Writer-", writers)
	}
}
