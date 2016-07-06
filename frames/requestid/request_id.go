package requestid

import (
	"net/http"

	"github.com/AlexanderChen1989/stack"
)

const (
	// DefaultHTTPHeader default http header for request id
	DefaultHTTPHeader = "x-request-id"
)

// New create a new request id Frame with customized http header
func New(header string) stack.FrameFunc {
	if header == "" {
		header = DefaultHTTPHeader
	}
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		setRequestID(header, w, getRequestID(header, r))
		next(w, r)
	}
}

func validRequestID(rid string) bool {
	size := len(rid)
	return size >= 20 && size <= 200
}

func genRequestID() string {
	return randString(32)
}

func getRequestID(header string, r *http.Request) string {
	rid := r.Header.Get(header)

	if validRequestID(rid) {
		return rid
	}

	return genRequestID()
}

func setRequestID(header string, w http.ResponseWriter, rid string) {
	w.Header().Add(header, rid)
}
