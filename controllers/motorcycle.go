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
	Categories  []string `json:"categories"`
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
	chain := db.Where("make != \"\" AND model != \"\"")
	for i, category := range body.Categories {
		query := "category LIKE ?"
		if i > 0 {
			chain = chain.Or(query, fmt.Sprintf("%%%s%%", category))
		} else {
			chain = chain.Where(query, fmt.Sprintf("%%%s%%", category))
		}
	}
	if body.Budget != 0 {
		chain = chain.Where("price <= ? AND price != 0", body.Budget)
	}
	if body.SeatHeight != 0 {
		chain = chain.Where("seat_height <= ?", body.SeatHeight)
	}
	if body.YearStart != 0 {
		chain = chain.Where("year_end >= ? OR year_end = 0", body.YearStart)
	}
	if body.YearEnd != math.MaxUint16 {
		chain = chain.Where("year_start <= ? AND year_start != 0", body.YearEnd)
	}
	if len(body.EngineTypes) > 0 {
		chain = chain.Where("engine_type IN ?", body.EngineTypes)
	}

	chain = chain.Order("overall_rating DESC")
	if len(body.OrderBy) > 0 {
		for _, order := range body.OrderBy {
			order = fmt.Sprintf("Review__%s DESC", order)
			chain = chain.Order(order)
		}
	}
	return chain
}

func GetMotorcycleRecommendations(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	topStr := queryParams.Get("top")
	top, err := strconv.ParseUint(topStr, 10, 16)
	if err != nil {
		top = 5
	}
	var body motoRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	db := models.GetDB()
	var motorcycles []models.Motorcycle
	body.cleanup()
	chain := buildQuery(&body, db)
	chain.Debug().Joins("Review").Limit(int(top)).Find(&motorcycles)

	RespondJSON(w, http.StatusOK, motorcycles)
}
