package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	//api handler call for current City using here geocode api
	currentCityLat, currentCityLng := getCityCoordinates(r.FormValue("currentCity"))

	//api handler call for destination City using here geocode api
	destinationCityLat, destinationCityLng := getCityCoordinates(r.FormValue("destinationCity"))
	RouteDuration, RouteLength := getRouteBetweenCoordinates(currentCityLat, currentCityLng, destinationCityLat, destinationCityLng)

	//Print out current route
	fmt.Println(fmt.Sprintf("Duration: %d hours, Length: %d kilometres", RouteDuration, RouteLength))

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
