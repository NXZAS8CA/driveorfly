package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method ist not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesful\n")
	//current location
	currentCity := r.FormValue("currentCity")
	currentCountry := r.FormValue("currentCountry")
	//destination location
	destinationCity := r.FormValue("destinationCity")
	destinationCountry := r.FormValue("destinationCountry")

	fmt.Fprintf(w, "Current Location: %s, %s\n", currentCity, currentCountry)

	fmt.Fprintf(w, "Destination: %s, %s\n", destinationCity, destinationCountry)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}