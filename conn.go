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

// Deadline delegate to Context.Deadline
func (conn Conn) Deadline() (deadline time.Time, ok bool) {
	return conn.Context.Deadline()
}

// Done delegate to Context.Done
func (conn Conn) Done() <-chan struct{} {
	return conn.Context.Done()
}

// Err delegate to Context.Err
func (conn Conn) Err() error {
	return conn.Context.Err()
}

// Value delegate to Context.Value
func (conn Conn) Value(key interface{}) interface{} {
	return conn.Context.Value(key)
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
