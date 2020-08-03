package chi

import (
	"net/http"

	chipkg "github.com/go-chi/chi"
)

// Chi ...
type Chi struct {
	router *chipkg.Mux
}

// New ...
func New() *Chi {
	return &Chi{
		router: chipkg.NewRouter(),
	}
}

// AddRoute ...
func (c *Chi) AddRoute(method string, path string, handler http.HandlerFunc) {
	c.router.Method(method, path, handler)
}

// GetRouter ...
func (c *Chi) GetRouter() http.Handler {
	return c.router
}
