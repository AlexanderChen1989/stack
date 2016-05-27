package mux

import (
	"net/http"

	"github.com/AlexanderChen1989/plug"
	"golang.org/x/net/context"
)

// Plug implements plug.Plugger.Plug
func (r *Router) Plug(plug.Plugger) plug.Plugger {
	return r
}

// HandleConn implements plug.Plugger.HandleConn
func (r *Router) HandleConn(conn plug.Conn) {
	r.ServeHTTP(conn.Context, conn.ResponseWriter, conn.Request)
}

// Dispatch dispatch request to  plug.Plugger
func (r *Router) Dispatch(path string, p plug.Plugger) *Route {
	return r.Handle(path, ToHandler(p))
}

// DispatchFunc dispatch request to func(conn plug.Conn)
func (r *Router) DispatchFunc(path string, f func(conn plug.Conn)) *Route {
	return r.Handle(path, ToHandlerFunc(f))
}

// ToHandler convert plug.Plugger to Handler
func ToHandler(p plug.Plugger) Handler {
	return HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		p.HandleConn(plug.Conn{
			Context:        ctx,
			ResponseWriter: w,
			Request:        r,
		})
	})
}

// ToHandlerFunc convert func(conn plug.Conn) to HandlerFunc
func ToHandlerFunc(f func(conn plug.Conn)) HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		f(plug.Conn{
			Context:        ctx,
			ResponseWriter: w,
			Request:        r,
		})
	}
}
