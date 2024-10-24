package routes

import (
	"github.com/gorilla/mux"
	"github.com/sabillahsakti/task-management/controllers/authcontroller"
	"github.com/sabillahsakti/task-management/controllers/taskcontroller"
	"github.com/sabillahsakti/task-management/middlewares"
)

func SetupRoutes(r *mux.Router) {

	// Auth routes
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/task", taskcontroller.Create).Methods("POST")
	api.HandleFunc("/task/{id}", taskcontroller.GetByID).Methods("GET")
	api.HandleFunc("/task", taskcontroller.GetByUser).Methods("GET")
	api.HandleFunc("/task/{id}", taskcontroller.Update).Methods("PUT")
	api.HandleFunc("/task/{id}", taskcontroller.Delete).Methods("DELETE")

	api.Use(middlewares.JWTMiddleware)

}
