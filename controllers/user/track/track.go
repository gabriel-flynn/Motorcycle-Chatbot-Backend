package track

import (
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/controllers"
	"github.com/gabriel-flynn/Track-Locator/models"
	"github.com/gabriel-flynn/Track-Locator/services"
	"net"
	"net/http"
)

func GetClosestTrack(w http.ResponseWriter, r *http.Request) {
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

	travelTime, err := services.TravelTimeToTrack(user.Location, user.ClosestTrack)
	fmt.Println(travelTime)
	if err != nil {
		controllers.RespondJSON(w, http.StatusInternalServerError, err)
	}

	responseJson := map[string]interface{} {
		"travel_time": travelTime,
		"track": user.ClosestTrack,
	}
	controllers.RespondJSON(w, http.StatusOK, responseJson)
}
