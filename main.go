package main

import (
	"fmt"
	"github.com/gabriel-flynn/Track-Locator/routers"
	"github.com/gabriel-flynn/Track-Locator/utils"
	"log"
	"net/http"
)


func main() {
	utils.InitGetIPDB()
	r := routers.SetupRouter()

	fmt.Println("Server is starting")
	log.Fatal(http.ListenAndServe(":8081", r))
}