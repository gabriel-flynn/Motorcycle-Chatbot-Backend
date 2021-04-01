package models

type User struct {
	Name           string       `json:"name"`
	IPAddress      string       `gorm:"primaryKey" json:"ip_address"`
	Location       *Location    `gorm:"constraint:OnDelete:CASCADE" json:"location;"`
	LocationId     uint         `json:"-"`
	ClosestTrack   *Track       `json:"closest_track"`
	ClosestTrackId uint         `json:"-"`
	Motorcycles    []Motorcycle `gorm:"many2many:user_motorcycles;constraint:OnDelete:CASCADE" json:"motorcycles"`
}
