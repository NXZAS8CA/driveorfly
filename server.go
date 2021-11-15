package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Struct for .json from first api call
type City struct {
	Embedded struct {
		CitySearchResults []struct {
			Links struct {
				CityItem struct {
					Href string `json:"href"`
				} `json:"city:item"`
			} `json:"_links"`
			MatchingAlternateNames []struct {
				Name string `json:"name"`
			} `json:"matching_alternate_names"`
			MatchingFullName string `json:"matching_full_name"`
		} `json:"city:search-results"`
	} `json:"_embedded"`
	Links struct {
		Curies []struct {
			Href      string `json:"href"`
			Name      string `json:"name"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Count int `json:"count"`
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	//current location api request
	response, err := http.Get("https://api.teleport.org/api/cities/?search=" + r.FormValue("currentCity"))
	responseData, err := ioutil.ReadAll(response.Body)
	//create struct to store .json from request
	var currentCity City
	err = json.Unmarshal(responseData, &currentCity)
	if err != nil {
		log.Fatal(err)
	}

	requestedURLcurrentCity := currentCity.Embedded.CitySearchResults[0].Links.CityItem.Href
	fmt.Println(requestedURLcurrentCity)
	//destination location api request
	response, err = http.Get("https://api.teleport.org/api/cities/?search=" + r.FormValue("destinationCity"))
	responseData, err = ioutil.ReadAll(response.Body)
	//create stuct to store .json from request
	var destinationCity City
	err = json.Unmarshal(responseData, &destinationCity)
	if err != nil {
		log.Fatal(err)
	}

	requestedURLdestinationCity := destinationCity.Embedded.CitySearchResults[0].Links.CityItem.Href
	fmt.Println(requestedURLdestinationCity)
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
