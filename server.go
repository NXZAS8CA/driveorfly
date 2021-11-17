package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"github.com/mkrou/geonames"
	//"github.com/mkrou/geonames/models"
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

type City2 struct {
	Links struct {
		CityAdmin1Division struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"city:admin1_division"`
		CityAlternateNames struct {
			Href string `json:"href"`
		} `json:"city:alternate-names"`
		CityCountry struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"city:country"`
		CityTimezone struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"city:timezone"`
		CityUrbanArea struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"city:urban_area"`
		Curies []struct {
			Href      string `json:"href"`
			Name      string `json:"name"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	FullName  string `json:"full_name"`
	GeonameID int    `json:"geoname_id"`
	Location  struct {
		Geohash string `json:"geohash"`
		Latlon  struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"latlon"`
	} `json:"location"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

//Function and struct should be in own file
func formHandler(w http.ResponseWriter, r *http.Request) {

	//current location api request
	response, err := http.Get("https://api.teleport.org/api/cities/?search=" + r.FormValue("currentCity"))
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//create struct to store .json from request
	var currentCity City
	err = json.Unmarshal(responseData, &currentCity)
	if err != nil {
		log.Fatal(err)
	}

	//second api call to get the coords of given city
	response, err = http.Get(currentCity.Embedded.CitySearchResults[0].Links.CityItem.Href)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var currentCity2 City2
	err = json.Unmarshal(responseData, &currentCity2)
	if err != nil {
		log.Fatal(err)
	}

	//Coords from current city
	currentCityLatitude := currentCity2.Location.Latlon.Latitude
	currentCityLongitude := currentCity2.Location.Latlon.Longitude

	//-----------------destination location api request------------------
	response, err = http.Get("https://api.teleport.org/api/cities/?search=" + r.FormValue("destinationCity"))
	if err != nil {
		log.Fatal(err)
	}

	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//create stuct to store .json from request
	var destinationCity City
	err = json.Unmarshal(responseData, &destinationCity)
	if err != nil {
		log.Fatal(err)
	}

	//second api call to get the coords of given city
	response, err = http.Get(destinationCity.Embedded.CitySearchResults[0].Links.CityItem.Href)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var destinationCity2 City2
	err = json.Unmarshal(responseData, &destinationCity2)
	if err != nil {
		log.Fatal(err)
	}

	//coords from destination city
	destinationCityLatitude := destinationCity2.Location.Latlon.Latitude
	destinationCityLongitude := destinationCity2.Location.Latlon.Longitude
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
