package handlers

import (
	"fmt"
	"github.com/aanciaes/smart-grow-api/config/database"
	"net/http"
)


func Hello(w http.ResponseWriter, r *http.Request) {

	_ := database.Conn.Connection

	var _, err = fmt.Fprintf(w, "Hello world from go lang!")
	if err != nil {
		_, _ = fmt.Fprintf(w, "An error occured: %d", err)
	}
}