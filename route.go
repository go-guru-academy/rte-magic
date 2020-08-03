package rtemagic

import (
	"net/http"
	"time"

	responsewriter "github.com/go-guru-academy/rte-magic/response-writer"
)

// Middleware ...
type Middleware func(http.Handler) http.Handler

// Route ...
type Route struct {
	// The request path
	path string

	// The request method
	method string

	// middleware can be used to modify requests and/or responses
	middleware []Middleware

	// input is custom input needed to handle the request
	// input is typically a struct pointer
	input interface{}

	// The request handler
	handler http.HandlerFunc
}

// first ...
func (rt *Route) first() http.HandlerFunc {
	next := rt.chainMiddleware(0)
	return func(w http.ResponseWriter, r *http.Request) {
		rw := responsewriter.New(w, rt.input)
		rw.SetRequestStart(time.Now().UnixNano())
		defer func() {
			rw.SetRequestEnd(time.Now().UnixNano())
		}()
		next.ServeHTTP(rw, r)
	}
}

// chainMiddleware ...
func (rt *Route) chainMiddleware(i int) http.Handler {
	if i == len(rt.middleware) {
		return rt.handler
	}
	middleware := rt.middleware[i]
	return middleware(rt.chainMiddleware(i + 1))
}
