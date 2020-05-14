package handler

import (
	"fmt"
	"net/http"
	"time"
)

// Handler  hello heandler
func Handler(w http.ResponseWriter, r http.Request) {

	cTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, cTime)
}
