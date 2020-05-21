package main

import (
	"fmt"
	"net/http"
	// "github.com/gorilla/context"
	// "github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
