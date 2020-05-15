package handler

import (
	"fmt"
	"net/http"
)

// ProjectHandler  project heandler
func ProjectHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		w.Write([]byte("post request"))
	} else {
		fmt.Fprintf(w, "project work")
	}
}
