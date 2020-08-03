package responsewriter

import (
	"net/http"
)

// ResponseWriter ...
type ResponseWriter struct {
	w            http.ResponseWriter
	input        interface{}
	requestStart int64
	requestEnd   int64
}

// New ...
func New(w http.ResponseWriter, input interface{}) *ResponseWriter {
	return &ResponseWriter{
		w:     w,
		input: input,
	}
}

// Write ...
func (rw *ResponseWriter) Write(data []byte) (int, error) {
	return rw.w.Write(data)
}

// Header ...
func (rw *ResponseWriter) Header() http.Header {
	return rw.w.Header()
}

// WriteHeader ...
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.w.WriteHeader(statusCode)
}

// SetRequestStart ...
func (rw *ResponseWriter) SetRequestStart(requestStart int64) {
	rw.requestStart = requestStart
}

// SetRequestEnd ...
func (rw *ResponseWriter) SetRequestEnd(requestEnd int64) {
	rw.requestEnd = requestEnd
}

// GetInput ...
func GetInput(w http.ResponseWriter) interface{} {
	rw, ok := w.(*ResponseWriter)
	if !ok {
		panic("rte-magic: invalid response writer")
	}
	return rw.input
}
