package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Name string
	Address string `gorm:"unique"`
	URL string
	Latitude float32
	Longitude float32
}
