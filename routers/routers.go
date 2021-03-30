package routers

import (
	"github.com/gabriel-flynn/Track-Locator/controllers"
	"github.com/gabriel-flynn/Track-Locator/controllers/user"
	"github.com/gabriel-flynn/Track-Locator/controllers/user/location"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user", user.NewUser).Methods("POST")
	myRouter.HandleFunc("/user", user.GetUser).Methods("GET")
	myRouter.HandleFunc("/user/location", location.UpdateLocation).Methods("PATCH")
	myRouter.HandleFunc("/motorcycles", controllers.GetMotorcycles).Methods("POST")
	return myRouter
}
