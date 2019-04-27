package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aanciaes/smart-grow-api/security"
	"net/http"
)

type loginForm struct {
	Username string `json:username`
	Password string `json:password`
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t loginForm
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	if security.CheckPasswordHash(t.Username, t.Password) {
		token, err := generateJwt()

		if err == nil {
			var _, err = fmt.Fprintf(w, token)
			if err != nil {
				_, _ = fmt.Fprintf(w, "An error occured: %d", err)
			}
		} else {
			_, _ = fmt.Fprintf(w, "An error occured: %d", err)
		}
	} else {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}
}

func AdminFunc (w http.ResponseWriter, r *http.Request) {
	var _, err = fmt.Fprintf(w, "admin")
	if err != nil {
		_, _ = fmt.Fprintf(w, "An error occured: %d", err)
	}
}