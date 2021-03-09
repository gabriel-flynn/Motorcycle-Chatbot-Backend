package main

import (
	"github.com/gabriel-flynn/Track-Locator/routers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := routers.SetupRouter()
	log.Fatal(http.ListenAndServe(":8081", r))
}