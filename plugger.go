package plug

import "net/http"

// Plugger is a layer of logic that process request,
// Plugger can stack together to be a pipeline,
// This pipeline can process request layer by layer
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

// Plug implements Plugger.Plug
func (h handler) Plug(Plugger) Plugger {
	return h
}

// Handle implements Plugger.Handle
func (h handler) Handle(conn Conn) {
	h.h.ServeHTTP(conn.ResponseWriter, conn.Request)
}

// ToPlugger convert http.Handler to Plugger
func ToPlugger(h http.Handler) Plugger {
	return handler{h: h}
}

// ToPluggerFunc convert a http.HandleFunc to a Plugger
func ToPluggerFunc(fun func(w http.ResponseWriter, r *http.Request)) Plugger {
	return ToPlugger(http.HandlerFunc(fun))
}
