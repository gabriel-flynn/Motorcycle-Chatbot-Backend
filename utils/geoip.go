package utils

import (
	"github.com/oschwald/geoip2-golang"
	"log"
)

var db *geoip2.Reader

// This will persist the DB in memory (~60MB -> would cause less I/O than having to load it each time the endpoint is called.
func InitGetIPDB() {
	var err error
	db, err = geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
}

func GetGeoIPDB() *geoip2.Reader {
	return db
}
