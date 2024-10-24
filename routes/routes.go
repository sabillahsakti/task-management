package routes

import (
	"github.com/gorilla/mux"
	"github.com/sabillahsakti/task-management/controllers/authcontroller"
)

func SetupRoutes(r *mux.Router) {

	// Auth routes
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	// r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

}
