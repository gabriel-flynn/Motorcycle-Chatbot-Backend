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
	name string
}

func NewLocation(w http.ResponseWriter, r *http.Request) {
	models.GetDB()
	ip := net.ParseIP(r.RemoteAddr)
	geoIPDB := utils.GetGeoIPDB()
	record, err := geoIPDB.City(ip)
	if err != nil {
		fmt.Println(err)
	}
	closestTrack := services.FindClosestTrack(record.Location.Longitude, record.Location.Latitude)
	location := &models.Location{
		ClosestTrack:   closestTrack,
		ClosestTrackId: closestTrack.ID,
		Latitude:       record.Location.Latitude,
		Longitude:      record.Location.Longitude,
	}

	var body requestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		fmt.Fprintln(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	fmt.Println(body.name)

	user := &models.User{
		Name:      body.name,
		IPAddress: r.RemoteAddr,
		Location:  location,
	}
	db := models.GetDB()
	db.Save(user)
	fmt.Fprintf(w, "%s", record.City.Names["en"])
}