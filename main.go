package main

import "net/http"
import "fmt"
import _ "net/http/pprof"
import "github.com/gorilla/mux"

func handle(router *mux.Router, method, url string, fn func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(url, fn).Methods(method)
}

func makeMux() *http.ServeMux {
	router := mux.NewRouter()
	handle(router, "GET", "/caps", handleGetDinos)
	handle(router, "OPTIONS", "/timelines", handleOptionsTimelines)
	handle(router, "GET", "/timelines", handleGetTimelines)
	handle(router, "POST", "/timelines", handlePostTimelines)
	handle(router, "PUT", "/timelines", handlePutTimelines)
	handle(router, "DELETE", "/timelines", handleDeleteTimelines)
	handle(router, "POST", "/trigger/{dino}/{sensor}", handleTrigger)

	mux := http.NewServeMux()
	mux.Handle("/", router)

	return mux
}

func main() {
	err := state.LoadConfig("dinos.conf")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	state.Init()
	state.RestoreTimelines("timelines.json")

	for _, d := range state.dinos {
		fmt.Printf("%v => %v \n", d.Dino.Name, d.gpio)
	}

	handler := makeMux()

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	err = http.ListenAndServe(":8080", handler)
	fmt.Printf("http server died because: %v", err)
}
