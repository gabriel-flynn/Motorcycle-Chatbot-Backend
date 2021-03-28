package config

import (
	"log"
	"os"
)

type config struct {
	GoogleKey            string
	GooglePlaceSearchUrl string
	GooglePlaceDetailsUrl string
}

var Config *config

func init() {
	placeSearch := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json"
	placeDetails := "https://maps.googleapis.com/maps/api/place/details/json"
	googleKey := os.Getenv("google_api_key")
	if googleKey == "" {
		log.Fatal("Missing an API key! Please set your bing api to an environment variable called bing_api_key.")
	}

	Config = &config{googleKey, placeSearch, placeDetails}
}
