package routers

import (
	"github.com/gorilla/mux"
)
import controllers 	"github.com/davdwhyte87/smh-api/controllers"

// GetUserRoutes ... returns all the routes for users
func GetUserRoutes(router *mux.Router) {
	router.HandleFunc("/user/signup", controllers.CreateUser).Methods("POST")
}
