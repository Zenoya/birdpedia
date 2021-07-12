package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function. Can declare methods
	r.HandleFunc("/hello", HomeHandler).Methods("GET")
	return r
}

func main() {
	r := NewRouter()
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Homepage Handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
