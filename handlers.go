package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendJSON(thingy interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-type: text/json")
	err := json.NewEncoder(w).Encode(thingy)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON encoding error (probably a server bug): %v", err),
			http.StatusInternalServerError)
	}
}

func handleGetDinos(w http.ResponseWriter, r *http.Request) {
	sendJSON(getCurrentDinoList())
}
