package plug

import (
	"net/http"
	"time"

	"golang.org/x/net/context"
)

// Conn group Context/Request/ResponseWriter together
type Conn struct {
	Context        context.Context
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

// Header delegate to ResponseWriter.Header
func (conn Conn) Header() http.Header {
	return conn.ResponseWriter.Header()
}

// Write delegate to ResponseWriter.Write
func (conn Conn) Write(data []byte) (int, error) {
	return conn.ResponseWriter.Write(data)
}

// WriteHeader delegate to ResponseWriter.WriteHeader
func (conn Conn) WriteHeader(code int) {
	conn.ResponseWriter.WriteHeader(code)
}

// WithCancel delegate to context.WithCancel
func WithCancel(conn Conn) (Conn, context.CancelFunc) {
	ctx, cancel := context.WithCancel(conn.Context)
	return Conn{
		Context:        ctx,
		Request:        conn.Request,
		ResponseWriter: conn.ResponseWriter,
	}, cancel
}

// WithDeadline delegate to context.WithDeadline
func WithDeadline(conn Conn, deadline time.Time) (Conn, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(conn.Context, deadline)
	return Conn{
		Context:        ctx,
		Request:        conn.Request,
		ResponseWriter: conn.ResponseWriter,
	}, cancel
}

// WithTimeout delegate to context.WithTimeout
func WithTimeout(conn Conn, timeout time.Duration) (Conn, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(conn.Context, timeout)
	return Conn{
		Context:        ctx,
		Request:        conn.Request,
		ResponseWriter: conn.ResponseWriter,
	}, cancel
}

// WithValue delegate to context.WithValue
func WithValue(conn Conn, key interface{}, val interface{}) Conn {
	ctx := context.WithValue(conn.Context, key, val)
	return Conn{
		Context:        ctx,
		Request:        conn.Request,
		ResponseWriter: conn.ResponseWriter,
	}
}
