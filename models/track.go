package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Name string
	City string
	State string
	Url string
	Latitude float32
	Longitude float32
}
