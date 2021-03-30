package routers

import (
	"github.com/gabriel-flynn/Track-Locator/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user", controllers.NewUser).Methods("POST")
	myRouter.HandleFunc("/user", controllers.GetUser).Methods("GET")
	myRouter.HandleFunc("/motorcycles", controllers.GetMotorcycles).Methods("POST")

	return myRouter
}
