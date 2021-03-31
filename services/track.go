package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/config"
	"github.com/gabriel-flynn/Track-Locator/models"
	"io/ioutil"
	"math"
	"net/http"
)

func haversine(lng1 float64, lat1 float64, lng2 float64, lat2 float64) float64 {
	latDist := (lat2 - lat1) * math.Pi / 180.0
	lngDist := (lng2 - lng1) * math.Pi / 180.0

	// convert latitudes to radians
	lat1 *= math.Pi / 180.0
	lat2 *= math.Pi / 180.0

	//haversine formula
	a := math.Pow(math.Sin(latDist/2), 2) +
		math.Pow(math.Sin(lngDist/2), 2)*
			math.Cos(lat1)*math.Cos(lat2)
	earthRadius := 6371.0
	c := 2 * math.Asin(math.Sqrt(a))
	return earthRadius * c
}

func FindClosestTrack(location *models.Location) (closestTrack *models.Track) {
	db := models.GetDB()
	var tracks []models.Track
	db.Find(&tracks)

	longitude := location.Longitude
	latitude := location.Latitude
	closestTrack = nil
	closestDist := math.MaxFloat64
	for _, track := range tracks {
		dist := haversine(track.Longitude, track.Latitude, longitude, latitude)
		if dist < closestDist {
			trackCopy := track
			closestTrack = &trackCopy
			closestDist = dist
		}
	}
	return
}

type distanceMatrixResponse struct {
	Rows []struct {
		Elements []struct {
			Duration struct {
				Text string `json:"text"`
			} `json:"duration"`
		} `json:"elements"`
	} `json:"rows"`
}

func TravelTimeToTrack(location *models.Location, track *models.Track) (string, error) {
	// Use Google Maps distance matrix API to calculate travel time
	request, _ := http.NewRequest("GET", config.Config.GoogleDistanceMatrixUrl, nil)

	//Query params
	q := request.URL.Query()
	q.Add("key", config.Config.GoogleKey)
	q.Add("origins", fmt.Sprintf("%f,%f", location.Latitude, location.Longitude))
	q.Add("destinations", fmt.Sprintf("%f,%f", track.Latitude, track.Longitude))
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return "", fmt.Errorf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var response distanceMatrixResponse
		json.Unmarshal(data, &response)
		if len(response.Rows) > 0 && len(response.Rows[0].Elements) > 0 {
			fmt.Println(response.Rows[0].Elements[0].Duration.Text)
			return response.Rows[0].Elements[0].Duration.Text, nil
		} else {
			return "", errors.New("could not find any places for that query")
		}
	}
}
