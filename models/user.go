package models

type User struct {
	Name string
	IPAddress string `gorm:"primaryKey"`
	Location *Location
	LocationId uint
}
