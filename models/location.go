package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	LocationString string  `json:"location_string"`
}
