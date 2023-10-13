package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
        return
	}

    if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
        return
    }

    fmt.Fprint(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return 
    }

    fmt.Fprintf(w, "POST successful!")
    name := r.FormValue("name")
    address := r.FormValue("address")
    fmt.Fprintf(w, "Name: %v, Address: %v", name, address)
}
