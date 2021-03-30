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

	fmt.Println("Server is listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
