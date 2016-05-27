package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexanderChen1989/plug"
	"github.com/stretchr/testify/assert"
)

func TestAddRequestID(t *testing.T) {
	p := newPlug(DefaultHTTPHeader)
	p.next = plug.NewEmpty()
	assert.Equal(t, p.httpHeader, DefaultHTTPHeader)
	req, _ := http.NewRequest("", "", nil)
	conn := plug.Conn{
		Request:        req,
		ResponseWriter: httptest.NewRecorder(),
	}
	fakeID := randString(32)
	conn.Request.Header.Add(DefaultHTTPHeader, fakeID)
	p.HandleConn(conn)
	assert.Equal(t, conn.ResponseWriter.Header().Get(DefaultHTTPHeader), fakeID)
	conn.Request.Header.Set(DefaultHTTPHeader, "")
	p.HandleConn(conn)
	assert.Len(t, conn.ResponseWriter.Header().Get(DefaultHTTPHeader), 32)
}
