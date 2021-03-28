package controllers

import (
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/models"
	"github.com/gabriel-flynn/Track-Locator/utils"
	"net"
	"net/http"
)



func NewLocation(w http.ResponseWriter, r *http.Request) {
	models.GetDB()
	//fmt.Fprintf(w, "This will be implemented eventually")
	ip := net.ParseIP(r.RemoteAddr)
	db := utils.GetGeoIPDB()
	record, err := db.City(ip)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "%s", record.City.Names["en"])
}