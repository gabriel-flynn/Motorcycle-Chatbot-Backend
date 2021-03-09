package routers

import (
	"github.com/gabriel-flynn/Track-Locator/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/location", controllers.NewLocation).Methods("POST")

	return myRouter
}