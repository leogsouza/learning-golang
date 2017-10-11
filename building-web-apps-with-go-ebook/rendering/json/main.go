package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:title`
	Author string `json:author`
}

func main() {
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":8080", nil) // I need to know what is the second parameter
}

func ShowBooks(rw http.ResponseWriter, r *http.Request) {
	books := []Book{
		Book{"Building Web Apps with Go", "Jeremy Saens"},
		Book{"Effective Go", "Jeremy Saens"},
	}

	JsonBooksEncoder(rw, books)
	JsonBooksMarshal(rw, books)

}

// Printing JSON with Encoder
func JsonBooksEncoder(rw http.ResponseWriter, books []Book) {
	encoder := json.NewEncoder(rw)
	encoder.SetIndent("  ", "    ")
	err := encoder.Encode(books)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func JsonBooksMarshal(rw http.ResponseWriter, books []Book) {
	js, err := json.Marshal(books)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Contenty-Type", "application/json")
	rw.Write(js)
}
