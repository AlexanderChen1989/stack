package plug

import "net/http"

type Plug interface {
	ServeHTTP(http.ResponseWriter, *http.Request, http.HandlerFunc)
}

type PlugFunc func(http.ResponseWriter, *http.Request, http.HandlerFunc)
