package models

type Location struct {
	Latitude float64 `gorm:"primaryKey; not null"`
	Longitude float64 `gorm:"primaryKey; not null"`
}
