package router

import (
	"net/http"

	"github.com/AlexanderChen1989/plug"
)

// New create a new Plug for router
func New(router http.Handler) plug.PlugFunc {
	if router == nil {
		panic("Router is nil")
	}

	return func(w http.ResponseWriter, r *http.Request, _ http.HandlerFunc) {
		router.ServeHTTP(w, r)
	}
}
