package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	// Router is the package level mux router.
	Router *mux.Router
)

// NewRouter creates a mux router, sets up
// a static handler and registers the dynamic
// routes.
func NewRouter() {
	Router = mux.NewRouter().StrictSlash(true)
	http.Handle("/", Router)
	Router.Methods("GET").Path("/").HandlerFunc(Root)
}

// Root is the hypermedia root endpoint of the service (GET).
func Root(w http.ResponseWriter, r *http.Request) {
	log.Fatal("BOOM")
}
