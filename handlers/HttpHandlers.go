package handlers

import (
	"fmt"
	"net/http"
)


func Login(w http.ResponseWriter, r *http.Request) {
	token, err := generateJwt()

	if err == nil {
		var _, err = fmt.Fprintf(w, token)
		if err != nil {
			_, _ = fmt.Fprintf(w, "An error occured: %d", err)
		}
	} else {
		_, _ = fmt.Fprintf(w, "An error occured: %d", err)
	}
}

func AdminFunc (w http.ResponseWriter, r *http.Request) {
	var _, err = fmt.Fprintf(w, "admin")
	if err != nil {
		_, _ = fmt.Fprintf(w, "An error occured: %d", err)
	}
}