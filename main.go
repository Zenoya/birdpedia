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

	// Declare the static file directory
	staticFileDirectory := http.Dir("./assets/")

	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

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
