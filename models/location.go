package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	name string
	closestTrack Track
	xCord float32
	yCord float32
}


//Will use the haversine formula to find closest coordinate https://github.com/umahmood/haversine
func (loc *Location) findClosestTrack(tracks *[]Track) {
	fmt.Println("Finding the closest track...")
	//do haversine formula on all the tracks, call bings maps API to get the closest track (factoring in driving distance) then return the closest one
}
