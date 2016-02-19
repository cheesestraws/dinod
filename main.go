package main

import "net/http"
import "fmt"
import "github.com/gorilla/mux"

func handle(router *mux.Router, method, url string, fn func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(url, fn).Methods(method)
}

func makeMux() *http.ServeMux {
	router := mux.NewRouter()
	handle(router, "GET", "/caps", handleGetDinos)

	mux := http.NewServeMux()
	mux.Handle("/", router)

	return mux
}

func main() {
	handler := makeMux()
	err := http.ListenAndServe(":8080", handler)
	fmt.Printf("http server died because: %v", err)
}
