package gplog

import (
	"fmt"
	"net/http"
)

// Trace trace request
func Trace(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := Logger(r)
	if logger == nil {
		fmt.Println("Please add log frame first")
	} else {
		defer logger.Trace("[Request]").End()
	}

	next(w, r)
}
