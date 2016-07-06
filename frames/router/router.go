package router

import (
	"net/http"

	"github.com/AlexanderChen1989/stack"
)

// New create a new Frame for router
func New(router http.Handler) stack.FrameFunc {
	if router == nil {
		panic("Router is nil")
	}

	return func(w http.ResponseWriter, r *http.Request, _ http.HandlerFunc) {
		router.ServeHTTP(w, r)
	}
}
