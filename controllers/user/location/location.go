package location

import (
	"encoding/json"
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/controllers"
	"github.com/gabriel-flynn/Track-Locator/models"
	"github.com/gabriel-flynn/Track-Locator/services"
	"net"
	"net/http"
)

type updateLocationRequestBody struct {
	LocationString string `json:"location_string"`
}

func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	ipStr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		//TODO: HANDLE ERROR
	}

	var user *models.User
	result := db.Joins("ClosestTrack").Joins("Location").First(&user, "ip_address = ?", ipStr)
	if result.Error != nil {
		var i struct{}
		controllers.RespondJSON(w, http.StatusNotFound, i)
	}

	var body updateLocationRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	user.Location.LocationString = body.LocationString
	services.SetLocationAndLatitude(user.Location)
	user.ClosestTrack = services.FindClosestTrack(user.Location)

	db.Save(user.Location)
	db.Save(user)

	controllers.RespondJSON(w, http.StatusOK, user)
}
