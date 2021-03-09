package main

import (
	"github.com/gabriel-flynn/Track-Locator/routers"
	"log"
	"net/http"
)


func main() {

	r := routers.SetupRouter()
	log.Fatal(http.ListenAndServe(":8081", r))
}