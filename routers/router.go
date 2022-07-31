package routers

import (
	"github.com/gorilla/mux"

	"goauth/controllers"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers).Methods("get")
	return router
}
