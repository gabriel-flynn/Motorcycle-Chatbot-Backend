package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Name      string  `json:"name"`
	Address   string  `gorm:"unique" json:"address"`
	URL       string  `json:"url"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
