package handlers

import (
	"fmt"
	"net/http"
)


func Hello(w http.ResponseWriter, r *http.Request) {
	var _, err = fmt.Fprintf(w, "Hello world from go lang!")
	if err != nil {
		_, _ = fmt.Fprintf(w, "An error occured: %d", err)
	}
}