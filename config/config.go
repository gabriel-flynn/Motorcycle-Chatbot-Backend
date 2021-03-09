package config

import (
	"fmt"
	"os"
)

type config struct {
	bingKey string
}

var Config *config

func init() {
	bingKey := os.Getenv("bing_api_key")
	if bingKey == "" {
		fmt.Println()
	}

	Config = &config{bingKey}
}
