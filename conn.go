package plug

import (
	"net/http"
	"time"

	"golang.org/x/net/context"
)

type Conn struct {
	Context        context.Context
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

func WithCancel(conn Conn) (Conn, context.CancelFunc) {
	ctx, cancel := context.WithCancel(conn.Context)
	return Conn{
		Context:        ctx,
		Request:        conn.Request,
		ResponseWriter: conn.ResponseWriter,
	}, cancel
}

func WithDeadline(conn Conn, deadline time.Time) (Conn, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(conn.Context, deadline)
	return Conn{
		Context:        ctx,
		Request:        conn.Request,
		ResponseWriter: conn.ResponseWriter,
	}, cancel
}

func WithTimeout(conn Conn, timeout time.Duration) (Conn, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(conn.Context, timeout)
	return Conn{
		Context:        ctx,
		Request:        conn.Request,
		ResponseWriter: conn.ResponseWriter,
	}, cancel
}

func WithValue(conn Conn, key interface{}, val interface{}) Conn {
	ctx := context.WithValue(conn.Context, key, val)
	return Conn{
		Context:        ctx,
		Request:        conn.Request,
		ResponseWriter: conn.ResponseWriter,
	}
}
