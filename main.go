package main

import (
	"crypto/tls"
	"fmt"
	"github.com/aanciaes/smart-grow-api/config/database"
	"github.com/aanciaes/smart-grow-api/handlers"
	"github.com/aanciaes/smart-grow-api/persistence"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"os"
)

func redirectHttpsHandler (w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func serveSingle(pattern string, filename string, r *mux.Router) {
	r.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func setHandlers (r *mux.Router) {
	r.Handle("/", http.FileServer(http.Dir("static")))
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.Handle("/temperature", handlers.SecureEndpoint(handlers.GetTemperature)).Methods("GET")
	r.Handle("/humidity", handlers.SecureEndpoint(handlers.GetHumidity)).Methods("GET")
	r.Handle("/light", handlers.SecureEndpoint(handlers.GetLight)).Methods("GET")
	r.Handle("/soil", handlers.SecureEndpoint(handlers.GetSoil)).Methods("GET")
	r.Handle("/temperature", handlers.SecureEndpoint(handlers.CreateTemperature)).Methods("POST")
	r.Handle("/humidity", handlers.SecureEndpoint(handlers.CreateHumidity)).Methods("POST")
	r.Handle("/light", handlers.SecureEndpoint(handlers.CreateLight)).Methods("POST")
	r.Handle("/soil", handlers.SecureEndpoint(handlers.CreateSoil)).Methods("POST")
	r.Handle("/artificialLight", handlers.SecureEndpoint(handlers.TurnOnOffLight)).Methods("POST")
	r.Handle("/waterPlants", handlers.SecureEndpoint(handlers.WaterPlants)).Methods("POST")

	r.Handle("/routine", handlers.SecureEndpoint(handlers.CreateRoutine)).Methods("POST")
	r.Handle("/routine", handlers.SecureEndpoint(handlers.GetRoutines)).Methods("GET")
	r.Handle("/routine", handlers.SecureEndpoint(handlers.DeleteRoutine)).Methods("DELETE")

	serveSingle("/favicon.ico", "./static/favicon.ico", r)
}

func getIpAddress () string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
	for _, a := range addr {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return "127.0.0.1"
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

		log.Printf("Starting SSL server, listening at %s:443\n", getIpAddress())

		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "dev"
		}

		// redirect every http request to https
		go http.ListenAndServe(":80", http.HandlerFunc(redirectHttpsHandler))
		go persistence.CheckRoutines()

		// Return database configuration based on environment variable
		if env == "prod" {
			log.Fatal(srv.ListenAndServeTLS("keys/fullchain.pem", "keys/privkey.pem"))
		} else {
			log.Fatal(srv.ListenAndServeTLS("keys/tls.crt", "keys/tls.key"))
		}
	} else {
		log.Fatal(err)
	}
}