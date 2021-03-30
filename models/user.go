package models

type User struct {
	Name           string    `json:"name"`
	IPAddress      string    `gorm:"primaryKey" json:"ip_address"`
	Location       *Location `json:"location"`
	LocationId     uint      `json:"-"`
	ClosestTrack   *Track    `json:"closest_track"`
	ClosestTrackId uint      `json:"-"`
}
