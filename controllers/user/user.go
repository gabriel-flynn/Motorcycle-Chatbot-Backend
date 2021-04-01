package user

import (
	"encoding/json"
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/controllers"
	"github.com/gabriel-flynn/Track-Locator/models"
	"github.com/gabriel-flynn/Track-Locator/services"
	"github.com/gabriel-flynn/Track-Locator/utils"
	"gorm.io/gorm/clause"
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

	city := record.City.Names["en"]
	state := ""
	if len(record.Subdivisions) > 0 {
		state = record.Subdivisions[0].Names["en"]
	}

	var locationStr string
	if city != "" && state != "" {
		locationStr = fmt.Sprintf("%s,%s", city, state)
	} else {
		locationStr = fmt.Sprintf("%s%s", city, state)
	}
	location := &models.Location{
		Latitude:       record.Location.Latitude,
		Longitude:      record.Location.Longitude,
		LocationString: locationStr,
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

	controllers.RespondJSON(w, http.StatusOK, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	ipStr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		//TODO: HANDLE ERROR
	}
	var user models.User
	result := db.Preload("Motorcycles.Review").Preload(clause.Associations).First(&user, "ip_address = ?", ipStr)
	if result.Error == nil {
		controllers.RespondJSON(w, http.StatusOK, user)
		return
	} else {
		var i struct{}
		controllers.RespondJSON(w, http.StatusNoContent, i)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	ipStr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		//TODO: HANDLE ERROR
	}
	var user models.User
	result := db.Delete(&user, "ip_address = ?", ipStr)
	if result.Error == nil {
		controllers.RespondJSON(w, http.StatusNoContent, nil)
		return
	} else {
		var i struct{}
		controllers.RespondJSON(w, http.StatusNotFound, i)
		return
	}
}
