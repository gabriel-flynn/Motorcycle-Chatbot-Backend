package models

import (
	"gorm.io/gorm"
)
//`gorm:"foreignkey:id;references:id;"`
type Location struct {
	gorm.Model
	Name string
	ClosestTrack *Track
	ClosestTrackId uint
	Latitude float64
	Longitude float64
}
