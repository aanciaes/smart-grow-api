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
	decoder.DisallowUnknownFields()

	var login loginForm
	err := decoder.Decode(&login)

	if err == nil {
		if security.CheckPasswordHash(login.Username, login.Password) {
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
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func AdminFunc(w http.ResponseWriter, r *http.Request) {
	var _, err = fmt.Fprintf(w, "admin")
	if err != nil {
		_, _ = fmt.Fprintf(w, "An error occured: %d", err)
	}
}
