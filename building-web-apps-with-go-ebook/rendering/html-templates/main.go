package main

import (
	"html/template"
	"net/http"
	"os"
	"path"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
