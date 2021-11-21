package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// City Struct for first api call in getCityCoordinates
type City struct {
	Items []struct {
		Title        string `json:"title"`
		ID           string `json:"id"`
		ResultType   string `json:"resultType"`
		LocalityType string `json:"localityType"`
		Address      struct {
			Label       string `json:"label"`
			CountryCode string `json:"countryCode"`
			CountryName string `json:"countryName"`
			StateCode   string `json:"stateCode"`
			State       string `json:"state"`
			CountyCode  string `json:"countyCode"`
			County      string `json:"county"`
			City        string `json:"city"`
			PostalCode  string `json:"postalCode"`
		} `json:"address"`
		Position struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"position"`
		MapView struct {
			West  float64 `json:"west"`
			South float64 `json:"south"`
			East  float64 `json:"east"`
			North float64 `json:"north"`
		} `json:"mapView"`
		Scoring struct {
			QueryScore float64 `json:"queryScore"`
			FieldScore struct {
				City float64 `json:"city"`
			} `json:"fieldScore"`
		} `json:"scoring"`
	} `json:"items"`
}

// Route struct for third api call in getRouteBetweenCoordinates
type Route struct {
	Routes []struct {
		ID       string `json:"id"`
		Sections []struct {
			ID        string `json:"id"`
			Type      string `json:"type"`
			Departure struct {
				Time  time.Time `json:"time"`
				Place struct {
					Type     string `json:"type"`
					Location struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"location"`
					OriginalLocation struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"originalLocation"`
				} `json:"place"`
			} `json:"departure"`
			Arrival struct {
				Time  time.Time `json:"time"`
				Place struct {
					Type     string `json:"type"`
					Location struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"location"`
					OriginalLocation struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"originalLocation"`
				} `json:"place"`
			} `json:"arrival"`
			Summary struct {
				Duration     int `json:"duration"`
				Length       int `json:"length"`
				BaseDuration int `json:"baseDuration"`
			} `json:"summary"`
			Transport struct {
				Mode string `json:"mode"`
			} `json:"transport"`
		} `json:"sections"`
	} `json:"routes"`
}

func getCityCoordinates(originCity string) (latitude float64, longitude float64) {

	response, err := http.Get(fmt.Sprintf("https://geocode.search.hereapi.com/v1/geocode?q=%s&apiKey=%s", originCity, os.Getenv("HERE_Routing_API")))
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

	return currentCity.Items[0].Position.Lat, currentCity.Items[0].Position.Lng
}
func getRouteBetweenCoordinates(originLatitude float64, originLongitude float64, destinationLatitude float64, destinationLongitude float64) (RouteDuration int, RouteLength int) {
	request := fmt.Sprintf("https://router.hereapi.com/v8/routes?transportMode=car&origin=%f,%f&destination=%f,%f&return=summary&apikey=%s", originLatitude, originLongitude, destinationLatitude, destinationLongitude, os.Getenv("HERE_Routing_API"))

	response, err := http.Get(request)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var route Route
	err = json.Unmarshal(responseData, &route)
	if err != nil {
		log.Fatal(err)
	}

	return route.Routes[0].Sections[0].Summary.Duration / 3600, route.Routes[0].Sections[0].Summary.Length / 1000
}
