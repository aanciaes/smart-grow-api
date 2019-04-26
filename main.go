package main

import (
	"github.com/aanciaes/smart-grow-api/config/database"
	"github.com/aanciaes/smart-grow-api/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setHandlers (r *mux.Router) {
	r.Handle("/", secureEndpoint(handlers.Hello))
}

func secureEndpoint (h http.HandlerFunc) http.Handler {
	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		h.ServeHTTP(w, r)
		defer log.Println("After")
	})
}

func main() {
	_, err := database.ConfigDatabase()
	database.BootstrapDatabase()

	r := mux.NewRouter()
	setHandlers(r)

	if err == nil {
		log.Println("Starting server, listening at port 8000")
		log.Fatal(http.ListenAndServe(":8000", r))
	} else {
		log.Fatal(err)
	}
}