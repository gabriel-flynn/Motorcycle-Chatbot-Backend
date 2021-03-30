package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Name      string
	Address   string `gorm:"unique"`
	URL       string
	Latitude  float64
	Longitude float64
}
