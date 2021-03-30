package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/models"
	"net/http"
	"strconv"
)

type motoRequestBody struct {
	Category    string   `json:"category"`
	Budget      float32  `json:"budget"`
	SeatHeight  uint8    `json:"seat_height"`
	YearStart   uint16   `json:"year_start"`
	YearEnd     uint16   `json:"year_end"`
	EngineTypes []string `json:"engine_types"`
	Extra       string   `json:"extra"`
}

func GetMotorcycles(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	topStr := queryParams.Get("top")
	top, err := strconv.ParseUint(topStr, 10, 16)
	if err != nil {
		top = 5
	}
	fmt.Println(top)
	fmt.Println("before body")
	var body motoRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request payload")
		return
	}
	fmt.Println("after body")
	defer r.Body.Close()
	db := models.GetDB()
	var motorcycles *models.Motorcycle
	db.First(&motorcycles)
	fmt.Println(motorcycles.Model + " " + motorcycles.EngineType)
	//ver motorcycles []*models
	//result := db.Find()
	fmt.Println(body.SeatHeight)
	fmt.Println(body.YearStart)
	fmt.Println(body.Budget)
	fmt.Println(body.Category)
	fmt.Println(body.EngineTypes[0])
	fmt.Println(body.EngineTypes[1])
	fmt.Println(body.Extra)
}
