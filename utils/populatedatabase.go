package utils

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/config"
	"github.com/gabriel-flynn/Track-Locator/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func PopulateTracksInDB() {
	csvFile, err := os.Open("tracks.csv")
	if err != nil {
		log.Fatalln("Could not open the csv file", err)
	}

	db := models.GetDB()
	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, rec := range record {
			query := strings.TrimSpace(rec)
			fmt.Println(query)
			placeId, err := getPlaceId(query)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Place ID: %s\n", placeId)
			}
			track, err := getPlaceDetails(placeId)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Track: %s\n\n", track)
			}
			db.Save(track)
		}
	}
}

type placeSearchResponse struct {
	Candidates []*placeID `json:"candidates"`
}

type placeID struct {
	Id string `json:"place_id"`
}

func getPlaceId(query string) (string, error) {
	request, _ := http.NewRequest("GET", config.Config.GooglePlaceSearchUrl, nil)

	//Query params
	q := request.URL.Query()
	q.Add("key", config.Config.GoogleKey)
	q.Add("inputtype", "textquery")
	q.Add("input", query)
	q.Add("fields", "place_id")
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var response placeSearchResponse
		json.Unmarshal(data, &response)
		if len(response.Candidates) > 0 {
			return response.Candidates[0].Id, nil
		} else {
			return "", errors.New("could not find any places for that query")
		}
	}
}

type placeDetailsResponse struct {
	Result struct {
		Address  string `json:"formatted_address"`
		Geometry struct {
			Location struct {
				Lat  float64 `json:"lat"`
				Long float64 `json:"lng"`
			}
		} `json:"Geometry"`
		Name    string `json:"name"`
		Website string `json:"website"`
	} `json:"result"`
	Status string `json:"status"`
}

func getPlaceDetails(placeId string) (*models.Track, error) {
	request, _ := http.NewRequest("GET", config.Config.GooglePlaceDetailsUrl, nil)

	//Query params
	q := request.URL.Query()
	q.Add("key", config.Config.GoogleKey)
	q.Add("place_id", placeId)
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var response placeDetailsResponse
		json.Unmarshal(data, &response)

		track := &models.Track{Name: response.Result.Name, Address: response.Result.Address, URL: response.Result.Website, Latitude: response.Result.Geometry.Location.Lat, Longitude: response.Result.Geometry.Location.Long}
		return track, nil
	}
}
