package cources

import (
	"fmt"
	"net/http"
)

func CourcesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "cources index handler")
}
