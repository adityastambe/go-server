package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprint(w, "Hello"); err != nil {
		errStr := fmt.Sprintf("Can't get page: %v", err)
		http.Error(w, errStr, http.StatusNotFound)
	}
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err :%v", err)
		return
	}
	if _, err := fmt.Fprintln(w, "POST request successful"); err != nil {
		http.Error(w, "POST request was not successful", http.StatusInternalServerError)
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	if _, err := fmt.Fprintf(w, "Name = %s\nAddress = %s", name, address); err != nil {
		http.Error(w, "Can't Get the Name and Address", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/", fileServer)
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("POST /form", formHandler)

	fmt.Println("Starting server at port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
