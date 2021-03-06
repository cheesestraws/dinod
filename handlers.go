package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func sendJSON(thingy interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	err := json.NewEncoder(w).Encode(thingy)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, fmt.Sprintf("JSON encoding error (probably a server bug): %v", err),
			http.StatusInternalServerError)
	}
}

func readJSON(thingyPtr interface{}, w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(thingyPtr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, fmt.Sprintf("JSON decoding error: %v", err),
			http.StatusInternalServerError)
		return err
	}

	if thingyPtr == nil {
		println("Invalid input JSON")
		http.Error(w, "Valid JSON but invalid data: You sent a null.  Please don't.",
			http.StatusInternalServerError)
		err := errors.New("object unmarshalled to nil")
		return err
	}

	return nil
}

func validateObject(v Valider, w http.ResponseWriter) error {
	err := v.Valid()
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

func handleOptionsTimelines(w http.ResponseWriter, r *http.Request) {
	sendJSON(struct{}{}, w)
}

func handleGetTimelines(w http.ResponseWriter, r *http.Request) {
	sendJSON(getCurrentTimelines(), w)
}

func handlePostTimelines(w http.ResponseWriter, r *http.Request) {
	var timelines []Timeline
	readJSON(&timelines, w, r)
	for _, t := range timelines {
		if validateObject(Valider(t), w) != nil {
			return
		}
	}

	sendJSON(replaceAllTimelines(timelines), w)
}

func handlePutTimelines(w http.ResponseWriter, r *http.Request) {
	var timeline Timeline
	readJSON(&timeline, w, r)
	if validateObject(Valider(timeline), w) != nil {
		return
	}

	sendJSON(addOrReplaceTimeline(timeline), w)
}

func handleDeleteTimelines(w http.ResponseWriter, r *http.Request) {
	deleteAllTimelines()
}

func handleTrigger(w http.ResponseWriter, r *http.Request) {
	dino := mux.Vars(r)["dino"]
	sensor := mux.Vars(r)["sensor"]
	trigger(dino, sensor)
}
