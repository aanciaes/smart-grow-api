package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aanciaes/smart-grow-api/model"
	"github.com/aanciaes/smart-grow-api/persistence"
	"github.com/aanciaes/smart-grow-api/security"
	"net/http"
)



func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var login model.LoginForm
	err := decoder.Decode(&login)

	if err == nil {
		user, authorized := security.CheckPasswordHash(login.Username, login.Password)
		if authorized {
			token, err := generateJwt(user)

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

func Register (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var register model.RegisterForm
	err := decoder.Decode(&register)

	if err == nil {
		if register.Password == register.ConfirmPassword {
			err := persistence.RegisterUser(register); if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				w.WriteHeader(201)
			}
		} else {
			http.Error(w, "Passwords don't match", http.StatusBadRequest)
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
