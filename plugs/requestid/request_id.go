package requestid

import "github.com/AlexanderChen1989/plug"

const (
	// DefaultHTTPHeader default http header for request id
	DefaultHTTPHeader = "x-request-id"
)

type requestIDPlug struct {
	next plug.Plugger

	httpHeader string
}

func newPlug(headers ...string) *requestIDPlug {
	header := DefaultHTTPHeader
	if len(headers) > 0 {
		header = headers[0]
	}
	return &requestIDPlug{httpHeader: header}
}

// New create a new request id Plugger
func New(headers ...string) plug.Plugger {
	return newPlug(headers...)
}

func (p *requestIDPlug) Plug(next plug.Plugger) plug.Plugger {
	p.next = next
	return p
}

func validRequestID(rid string) bool {
	size := len(rid)
	return size >= 20 && size <= 200
}

func genRequestID() string {
	return randString(32)
}

func (p *requestIDPlug) getRequestID(conn plug.Conn) string {
	rid := conn.Request.Header.Get(p.httpHeader)

	if validRequestID(rid) {
		return rid
	}

	return genRequestID()
}

func (p *requestIDPlug) setRequestID(conn plug.Conn, rid string) plug.Conn {
	conn.ResponseWriter.Header().Add(p.httpHeader, rid)
	return conn
}

func (p *requestIDPlug) Handle(conn plug.Conn) {
	rid := p.getRequestID(conn)
	conn = p.setRequestID(conn, rid)
	p.next.Handle(conn)
}
