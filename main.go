package main

import (
	"crypto/tls"
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
		cfg := &tls.Config{
			MinVersion:               tls.VersionTLS11,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,

				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			},
		}

		srv := &http.Server{
			Addr:         ":443",
			Handler:      handlers.LoggingWrapper(os.Stdout, r),
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}

		log.Println("Starting SSL server, listening at port 443")
		log.Fatal(srv.ListenAndServeTLS("keys/tls.crt", "keys/tls.key"))
	} else {
		log.Fatal(err)
	}
}