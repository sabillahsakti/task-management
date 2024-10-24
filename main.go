package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabillahsakti/task-management/config"
	"github.com/sabillahsakti/task-management/routes"
)

func main() {

	config.ConnectDatabase()

	r := mux.NewRouter()

	routes.SetupRoutes(r)

	//Start the server
	log.Println("Server Berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
