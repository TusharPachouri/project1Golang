package main

import (
	"fmt"
	"log"
	"net/http"
)

func formsHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "post request Successful\n")
	Name := r.FormValue("Naam")
	Address := r.FormValue("pata")
	fmt.Fprintf(w, "Name: %s\n", Name)
	fmt.Fprintf(w, "Address: %s\n", Address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found ", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not found", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formsHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Start server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
