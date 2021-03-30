package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/config"
	"github.com/gabriel-flynn/Track-Locator/models"
	"io/ioutil"
	"net/http"
)

func SetLocationAndLatitude(location *models.Location) error {
	query := ""
	if location.City != "" {
		query += location.City
	}
	if location.State != "" {
		query += ", " + location.State
	}

	lat, long, err := makePlacesApiCall(query)
	if err != nil {
		return err
	}
	location.Latitude = lat
	location.Longitude = long

	return nil
}

type placesApiResponse struct {
	Candidates []struct {
		Geometry struct {
			Location struct {
				Latitude  float64 `json:"lat"`
				Longitude float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"candidates"`
}

func makePlacesApiCall(query string) (float64, float64, error) {
	request, _ := http.NewRequest("GET", config.Config.GooglePlaceSearchUrl, nil)

	//Query params
	q := request.URL.Query()
	q.Add("key", config.Config.GoogleKey)
	q.Add("inputtype", "textquery")
	q.Add("input", query)
	q.Add("fields", "geometry")
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return 0, 0, fmt.Errorf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var response placesApiResponse
		json.Unmarshal(data, &response)
		if len(response.Candidates) > 0 {
			return response.Candidates[0].Geometry.Location.Latitude, response.Candidates[0].Geometry.Location.Longitude, nil
		} else {
			return 0, 0, errors.New("could not find any places for that query")
		}

	}
}
