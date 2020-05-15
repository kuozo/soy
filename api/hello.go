package handler

import (
	"fmt"
	"net/http"

	"github.com/kuozo/soy/pkg"
)

// Handler  hello heandler
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, pkg.GetCurrentTime())
}
