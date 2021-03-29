package models

type User struct {
	Name string
	IPAddress string `gorm:"primaryKey"`
	Location *Location `gorm:"ForeignKey:Latitude,Longitude;"`
	Latitude float64 `json:"-"`
	Longitude float64 `json:"-"`
	ClosestTrack *Track
	ClosestTrackId uint `json:"-"`
}
