package main

import (
	"fmt"
	"net/http"
	"os"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
