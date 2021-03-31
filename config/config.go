package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	GoogleKey               string
	GooglePlaceSearchUrl    string
	GooglePlaceDetailsUrl   string
	GoogleDistanceMatrixUrl string
}

var Config *config

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	placeSearch := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json"
	placeDetails := "https://maps.googleapis.com/maps/api/place/details/json"
	distanceMatrix := "https://maps.googleapis.com/maps/api/distancematrix/json"
	googleKey := os.Getenv("google_api_key")
	if googleKey == "" {
		log.Fatal("Missing an API key! Please set your goggle api key to an environment variable called google_api_key.")
	}

	Config = &config{googleKey, placeSearch, placeDetails, distanceMatrix}
}
