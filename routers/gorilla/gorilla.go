package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Gorilla ...
type Gorilla struct {
	router *mux.Router
}

// New ...
func New() *Gorilla {
	return &Gorilla{
		router: mux.NewRouter(),
	}
}

// AddRoute ...
func (g *Gorilla) AddRoute(method string, path string, handler http.HandlerFunc) {
	g.router.HandleFunc(path, handler).Methods(method)
}

// GetRouter ...
func (g *Gorilla) GetRouter() http.Handler {
	return g.router
}
