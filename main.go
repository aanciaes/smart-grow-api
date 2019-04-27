package main

import (
	"github.com/aanciaes/smart-grow-api/config/database"
	"github.com/aanciaes/smart-grow-api/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func setHandlers (r *mux.Router) {
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.Handle("/admin", handlers.SecureEndpoint(handlers.AdminFunc)).Methods("GET")
}

func main() {
	_, err := database.ConfigDatabase()
	database.BootstrapDatabase()

	r := mux.NewRouter()
	setHandlers(r)

	if err == nil {
		log.Println("Starting server, listening at port 8000")
		log.Fatal(http.ListenAndServe(":8000", handlers.LoggingWrapper(os.Stdout, r)))
	} else {
		log.Fatal(err)
	}
}