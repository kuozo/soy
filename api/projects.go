package handler

import (
	"fmt"
	"net/http"
)

// ProjectHandler  project heandler
func ProjectHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "post" {
		w.Write([]byte("post request"))
	}
	fmt.Fprintf(w, "project work")
}
