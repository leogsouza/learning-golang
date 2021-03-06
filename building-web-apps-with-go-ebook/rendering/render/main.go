package main

import (
	"net/http"
	"os"

	"gopkg.in/unrolled/render.v1"
)

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, visit sub pages now."))
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Some binary data here."))
	})

	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"Hello": "json"})
	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		// Assumes you have a template in ./tamplates called "example.tmpl"
		// $mkdir -p template && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl

		r.HTML(w, http.StatusOK, "example", nil)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
