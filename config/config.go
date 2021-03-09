package config

import (
	"log"
	"os"
)

type config struct {
	BingKey string
}

var Config *config

func init() {
	bingKey := os.Getenv("bing_api_key")
	if bingKey == "" {
		log.Fatal("Missing an API key! Please set your bing api to an environment variable called bing_api_key.")
	}

	Config = &config{bingKey}
}
