package handler

import (
	"fmt"
	"net/http"
)

// IndexHandler Index Handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello index api.")
}
