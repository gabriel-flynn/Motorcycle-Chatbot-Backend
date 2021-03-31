package routers

import (
	"github.com/gabriel-flynn/Track-Locator/controllers"
	"github.com/gabriel-flynn/Track-Locator/controllers/user"
	"github.com/gabriel-flynn/Track-Locator/controllers/user/location"
	"github.com/gabriel-flynn/Track-Locator/controllers/user/motorcycles"
	"github.com/gabriel-flynn/Track-Locator/controllers/user/track"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user", user.NewUser).Methods("POST")
	myRouter.HandleFunc("/user", user.GetUser).Methods("GET")
	myRouter.HandleFunc("/user/location", location.UpdateLocation).Methods("PATCH")
	myRouter.HandleFunc("/user/track", track.GetClosestTrack).Queries("closest", "true").Methods("GET")
	myRouter.HandleFunc("/user/motorcycles", motorcycles.SaveMotorcycles).Methods("POST")
	myRouter.HandleFunc("/motorcycles", controllers.GetMotorcycleRecommendations).Methods("POST")
	return myRouter
}
