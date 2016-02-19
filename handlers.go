package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func sendJSON(thingy interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-type", "text/json")
	err := json.NewEncoder(w).Encode(thingy)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON encoding error (probably a server bug): %v", err),
			http.StatusInternalServerError)
	}
}

func readJSON(thingyPtr *Valider, w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(thingyPtr)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON decoding error: %v", err),
			http.StatusInternalServerError)
		return err
	}

	if thingyPtr == nil {
		http.Error(w, "Valid JSON but invalid data: You sent a null.  Please don't.",
			http.StatusInternalServerError)
		err := errors.New("object unmarshalled to nil")
		return err
	}

	err = (*thingyPtr).Valid()
	if err != nil {
		http.Error(w, fmt.Sprintf("Valid JSON but invalid data: %v", err),
			http.StatusInternalServerError)
		return err
	}

	return nil
}

func handleGetDinos(w http.ResponseWriter, r *http.Request) {
	sendJSON(getCurrentDinoList(), w)
}

func handleGetTimelines(w http.ResponseWriter, r *http.Request) {
	sendJSON(getCurrentTimelines(), w)
}

func handlePostTimelines(w http.ResponseWriter, r *http.Request) {
	var timeline Valider = Timeline{}
	readJSON(&timeline, w, r)
	str := fmt.Sprintf("%v", timeline)
	print(str)
}

func handlePutTimelines(w http.ResponseWriter, r *http.Request) {

}

func handleDeleteTimeLines(w http.ResponseWriter, r *http.Request) {

}
