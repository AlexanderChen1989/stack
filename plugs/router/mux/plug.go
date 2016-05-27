package mux

import (
	"net/http"

	"github.com/AlexanderChen1989/plug"
	"golang.org/x/net/context"
)

func (r *Router) Plug(plug.Plugger) plug.Plugger {
	return r
}

func (r *Router) HandleConn(conn plug.Conn) {
	r.ServeHTTP(conn.Context, conn.ResponseWriter, conn.Request)
}

func ToHandler(p plug.Plugger) Handler {
	return HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		p.HandleConn(plug.Conn{
			Context:        ctx,
			ResponseWriter: w,
			Request:        r,
		})
	})
}
