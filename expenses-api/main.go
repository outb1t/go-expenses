package main

import (
	"encoding/json"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rates)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}

func rates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(getCurrencies())
}
