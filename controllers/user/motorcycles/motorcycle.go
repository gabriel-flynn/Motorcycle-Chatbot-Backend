package motorcycles

import (
	"encoding/json"
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/controllers"
	"github.com/gabriel-flynn/Track-Locator/models"
	"net"
	"net/http"
)

func SaveMotorcycles(w http.ResponseWriter, r *http.Request) {

	var motorcycles []models.Motorcycle
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&motorcycles); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	db := models.GetDB()
	ipStr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		//TODO: HANDLE ERROR
	}

	fmt.Println(motorcycles[0].Review.Id)
	var user models.User
	result := db.Preload("Motorcycles").First(&user, "ip_address = ?", ipStr)
	if result.Error != nil {
		var i struct{}
		controllers.RespondJSON(w, http.StatusNotFound, i)
	} else {
		user.Motorcycles = motorcycles
		db.Exec("DELETE FROM `track-locator`.user_motorcycles WHERE user_ip_address = ?", ipStr)
		db.Model(&user).Association("Motorcycles").Clear()
		db.Model(&user).Association("Motorcycles").Append(&motorcycles)
	}
	controllers.RespondJSON(w, http.StatusOK, motorcycles)
}
