package rmux

import (
	"net/http"

	"github.com/AlexanderChen1989/plug"
	m "github.com/gorilla/mux"
)

// New create a new Plug for *mux.Router
func New(router *m.Router) plug.PlugFunc {
	if router == nil {
		panic("Router is nil")
	}
	return func(w http.ResponseWriter, r *http.Request, _ http.HandlerFunc) {
		router.ServeHTTP(w, r)
	}
}
