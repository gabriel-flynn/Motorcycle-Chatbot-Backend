package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
	State     string  `json:"state"`
}
