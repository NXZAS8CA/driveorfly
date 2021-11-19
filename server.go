package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	//api handler call for current city
	originCityLatitude, originCityLongitude := getCityCoordinations(getCityInformationURL(r.FormValue("currentCity")))
	fmt.Println(originCityLatitude, originCityLongitude)

	//api handler call for destination city
	destinationCityLatitude, destinationCityLongitude := getCityCoordinations(getCityInformationURL(r.FormValue("destinationCity")))
	fmt.Println(destinationCityLatitude, destinationCityLongitude)

	request := fmt.Sprintf("https://router.hereapi.com/v8/routes?transportMode=car&origin=%f,%f&destination=%f,%f&return=summary", originCityLatitude, originCityLongitude, destinationCityLatitude, destinationCityLongitude)
	fmt.Println(request)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Println(os.Getenv("HERE_Routing_API"))

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
