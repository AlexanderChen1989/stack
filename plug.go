package plug

import "net/http"

// Plug is a middleware
type Plug interface {
	ServeHTTP(http.ResponseWriter, *http.Request, http.HandlerFunc)
}

// PlugFunc is middleware func
type PlugFunc func(http.ResponseWriter, *http.Request, http.HandlerFunc)
