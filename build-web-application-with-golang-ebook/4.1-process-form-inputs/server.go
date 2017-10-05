package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello Leonardo!") // send data to client side - write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		username := template.HTMLEscapeString(r.Form.Get("username"))
		password := template.HTMLEscapeString(r.Form.Get("password"))
		fmt.Println("username", username, "password", password)
		if username == "leonardo" && password == "leo123" {
			t, _ := template.ParseFiles("upload.gtpl")
			t.Execute(w, nil)
		} else {
			t, _ := template.ParseFiles("login.gtpl")
			t.Execute(w, nil)
		}

	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadFile")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) // setting router rule
	http.HandleFunc("/login", login)   // set router
	http.HandleFunc("/upload", upload)
	fmt.Println(os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
