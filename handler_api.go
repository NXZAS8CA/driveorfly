package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Struct for first api call in getCityInformation
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

//struct for second api call in getCityCoordinations
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

func getCityInformationURL(input string) (output string) {
	response, err := http.Get("https://api.teleport.org/api/cities/?search=" + input)
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
	return currentCity.Embedded.CitySearchResults[0].Links.CityItem.Href
}

func getCityCoordinations(input string) (latitude float64, longitude float64) {
	response, err := http.Get(input)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var currentCity2 City2
	err = json.Unmarshal(responseData, &currentCity2)
	if err != nil {
		log.Fatal(err)
	}
	return currentCity2.Location.Latlon.Latitude, currentCity2.Location.Latlon.Longitude

}
