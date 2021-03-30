package services

import (
	"github.com/gabriel-flynn/Track-Locator/models"
	"math"
)

func haversine(lng1 float64, lat1 float64, lng2 float64, lat2 float64) float64 {
	latDist := (lat2 - lat1) * math.Pi / 180.0
	lngDist := (lng2 - lng1) * math.Pi / 180.0

	// convert latitudes to radians
	lat1 *= math.Pi / 180.0
	lat2 *= math.Pi / 180.0

	//haversine formula
	a := math.Pow(math.Sin(latDist/2), 2) +
		math.Pow(math.Sin(lngDist/2), 2)*
			math.Cos(lat1)*math.Cos(lat2)
	earthRadius := 6371.0
	c := 2 * math.Asin(math.Sqrt(a))
	return earthRadius * c
}

func FindClosestTrack(location *models.Location) (closestTrack *models.Track) {
	db := models.GetDB()
	var tracks []models.Track
	db.Find(&tracks)

	longitude := location.Longitude
	latitude := location.Latitude
	closestTrack = nil
	closestDist := math.MaxFloat64
	for _, track := range tracks {
		dist := haversine(track.Longitude, track.Latitude, longitude, latitude)
		if dist < closestDist {
			trackCopy := track
			closestTrack = &trackCopy
			closestDist = dist
		}
	}
	return
}

func TravelTimeToTrack(location *models.Location) {
	// Use Google Maps distance matrix API to calculate travel time
}
