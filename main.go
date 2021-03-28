package main

import (
	"github.com/gabriel-flynn/Track-Locator/routers"
	"github.com/gabriel-flynn/Track-Locator/utils"
	"log"
	"net/http"
)


func main() {
	utils.InitGetIPDB()
	r := routers.SetupRouter()
	log.Fatal(http.ListenAndServe(":8081", r))
}