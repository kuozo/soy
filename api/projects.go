package handler

import (
	"fmt"
	"net/http"
)

// ProjectHandler  project heandler
func ProjectHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "project work")
}
