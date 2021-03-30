package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/models"
	"gorm.io/gorm"
	"math"
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
	OrderBy     []string `json:"order_by"`
}

func (r *motoRequestBody) cleanup() {
	//Handle years
	if r.YearStart <= 1800 {
		r.YearStart = 0
	}
	if r.YearEnd <= 1800 {
		r.YearEnd = math.MaxUint16
	}

}

func buildQuery(body *motoRequestBody, db *gorm.DB) *gorm.DB {
	//Need to clean up the database -> looks of missing info
	db = db.Where("make != \"\" AND model != \"\"")
	if body.Category != "" {
		db = db.Where("category LIKE %?%", body.Category)
	}
	if body.Budget != 0 {
		db = db.Where("price <= ? AND price != 0", body.Budget)
	}
	if body.SeatHeight != 0 {
		db = db.Where("seat_height <= ?", body.SeatHeight)
	}
	if body.YearStart != 0 {
		db = db.Where("year_start >= ?", body.YearStart)
	}
	if body.YearEnd != math.MaxUint16 {
		db = db.Where("year_end <= ? AND year_end != 0", body.YearEnd)
	}
	if len(body.EngineTypes) > 0 {
		db = db.Where("engine_type IN ?", body.EngineTypes)
	}

	db = db.Order("overall_rating")
	if len(body.OrderBy) > 0 {
		for _, orderValue := range body.OrderBy {
			db = db.Order(orderValue)
		}
	}
	return db
}

func GetMotorcycles(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	topStr := queryParams.Get("top")
	top, err := strconv.ParseUint(topStr, 10, 16)
	if err != nil {
		top = 5
	}
	fmt.Println(top)
	var body motoRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	db := models.GetDB()
	var motorcycles []*models.Motorcycle
	body.cleanup()
	db = buildQuery(&body, db)
	db.Debug().Joins("Review").Find(&motorcycles)

	//for _, motorcycle := range motorcycles {
	//	fmt.Printf("ID: %d Make: %s Engine Type: %s, YearStart: %d YearEnd: %d", motorcycle.Id, motorcycle.Make, motorcycle.EngineType, motorcycle.YearStart, motorcycle.YearEnd)
	//}
	//fmt.Println(result.Statement)
	//fmt.Println(motorcycles.Model + " " + motorcycles.EngineType)
	//ver motorcycles []*models
	//result := db.Find()

	fmt.Println(body.SeatHeight)
	fmt.Println(body.YearStart)
	fmt.Println(body.Budget)
	fmt.Println(body.Category)
	fmt.Println(body.EngineTypes[0])
}
