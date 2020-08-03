package rtemagic

import (
	"net/http"

	responsewriter "github.com/go-guru-academy/rte-magic/response-writer"
	"github.com/go-guru-academy/rte-magic/routers/chi"
	"github.com/go-guru-academy/rte-magic/routers/gorilla"
)

const (
	GORILLA int = iota + 1 // Start at 1 to error on no input
	CHI
)

// Rte ...
type Rte struct {
	router Router
}

// Router ...
type Router interface {
	AddRoute(method string, path string, handler http.HandlerFunc)
	GetRouter() http.Handler
}

// New ...
func New(routerPkg int) *Rte {
	var router Router
	switch routerPkg {
	case GORILLA:
		router = gorilla.New()
	case CHI:
		router = chi.New()
	default:
		panic("rte-magic: invalid router type")
	}
	return &Rte{
		router: router,
	}
}

// GetInput ...
func GetInput(w http.ResponseWriter) interface{} {
	return responsewriter.GetInput(w)
}

// ServeHTTP ...
func (r *Rte) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	r.router.GetRouter().ServeHTTP(responseWriter, request)
}

// addRoute ...
func (r *Rte) addRoute(route *Route) {
	r.router.AddRoute(route.method, route.path, route.first())
}

// Post ...
func (r *Rte) Post(path string, handler http.HandlerFunc, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodPost,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

// Get ...
func (r *Rte) Get(path string, handler http.HandlerFunc, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodGet,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

// Put ...
func (r *Rte) Put(path string, handler http.HandlerFunc, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodPut,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

// Delete ...
func (r *Rte) Delete(path string, handler http.HandlerFunc, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodDelete,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

// Connect ...
func (r *Rte) Connect(path string, handler http.HandlerFunc, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodConnect,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

// Head ...
func (r *Rte) Head(path string, handler http.HandlerFunc, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodHead,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

// Options ...
func (r *Rte) Options(path string, handler http.HandlerFunc, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodOptions,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

// Patch ...
func (r *Rte) Patch(path string, handler http.HandlerFunc, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodPatch,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

// Trace ...
func (r *Rte) Trace(path string, handler http.HandlerFunc, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodTrace,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}
