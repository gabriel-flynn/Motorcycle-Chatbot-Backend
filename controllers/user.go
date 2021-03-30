package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/models"
	"github.com/gabriel-flynn/Track-Locator/services"
	"github.com/gabriel-flynn/Track-Locator/utils"
	"net"
	"net/http"
)

type requestBody struct {
	Name string
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	ipStr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		//TODO: HANDLE ERROR
	}
	ip := net.ParseIP(ipStr)
	geoIPDB := utils.GetGeoIPDB()
	record, err := geoIPDB.City(ip)
	if err != nil {
		fmt.Println(err)
	}
	location := &models.Location{
		Latitude:  record.Location.Latitude,
		Longitude: record.Location.Longitude,
	}
	closestTrack := services.FindClosestTrack(location)
	var body requestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	user := &models.User{
		Name:           body.Name,
		IPAddress:      ipStr,
		ClosestTrack:   closestTrack,
		ClosestTrackId: closestTrack.ID,
		Location:       location,
	}
	db := models.GetDB()
	db.Save(user)

	respondJSON(w, http.StatusOK, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	ipStr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		//TODO: HANDLE ERROR
	}
	var user *models.User
	result := db.Joins("ClosestTrack").Joins("Location").First(&user, "ip_address = ?", ipStr)

	if result.Error == nil {
		respondJSON(w, http.StatusOK, user)
	} else {
		var i struct{}
		respondJSON(w, http.StatusNoContent, i)
	}
}
