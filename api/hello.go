package handler

import (
	"fmt"
	"net/http"

	"github.com/kuozo/gopkg"
)

// Handler  hello heandler
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, gopkg.FormatNow())
}
