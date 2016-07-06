package stack

import "net/http"

// Frame is a middleware
type Frame interface {
	ServeHTTP(http.ResponseWriter, *http.Request, http.HandlerFunc)
}

// FrameFunc is middleware func
type FrameFunc func(http.ResponseWriter, *http.Request, http.HandlerFunc)
