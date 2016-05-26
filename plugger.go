package plug

import "net/http"

type Plugger interface {
	Plug(Plugger) Plugger
	Handle(Conn)
}

// HandleFunc handle function to process request
type HandleFunc func(Conn)

// Plug implement Plugger.Plug
func (f HandleFunc) Plug(Plugger) Plugger {
	return f
}

// Handle implement Plugger.Handle
func (f HandleFunc) Handle(conn Conn) {
	f(conn)
}

type handler struct {
	h http.Handler
}

func (h handler) Plug(Plugger) Plugger {
	return h
}

func (h handler) Handle(conn Conn) {
	h.h.ServeHTTP(conn.ResponseWriter, conn.Request)
}

func ToPlugger(h http.Handler) Plugger {
	return handler{h: h}
}

func ToPluggerFunc(fun func(w http.ResponseWriter, r *http.Request)) Plugger {
	return ToPlugger(http.HandlerFunc(fun))
}
