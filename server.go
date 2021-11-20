package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	//api handler call for current city
	originCityLatitude, originCityLongitude := getCityCoordinations(getCityInformationURL(r.FormValue("currentCity")))
	fmt.Println(originCityLatitude, originCityLongitude)

	//api handler call for destination city
	destinationCityLatitude, destinationCityLongitude := getCityCoordinations(getCityInformationURL(r.FormValue("destinationCity")))
	fmt.Println(destinationCityLatitude, destinationCityLongitude)
	fmt.Println(getRouteBetweenCoordinates(originCityLatitude, originCityLongitude, destinationCityLatitude, destinationCityLongitude))

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
