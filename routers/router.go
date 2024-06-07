package routers

import (
	"github/chrislcontrol/go-simple-api/views"

	"github.com/gorilla/mux"
)

func RouterV1() *mux.Router {
	API_V1 := "/api/v1"
	router := mux.NewRouter()
	router.HandleFunc(API_V1+"/home", views.Home).Methods("GET")
	router.HandleFunc(API_V1+"/products", views.Products).Methods("POST")

	return router
}
