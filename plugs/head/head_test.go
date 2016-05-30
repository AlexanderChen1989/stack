package head

import (
	"net/http"
	"testing"

	"github.com/AlexanderChen1989/plug"
	"github.com/stretchr/testify/assert"
)

func TestHeadPlugger(t *testing.T) {
	testCases := []struct {
		method   string
		expected string
	}{
		{"GET", "GET"},
		{"POST", "POST"},
		{"head", "GET"},
		{"Head", "GET"},
		{"HEAD", "GET"},
	}

	for _, c := range testCases {
		req, _ := http.NewRequest(c.method, "/", nil)

		b := plug.NewBuilder()
		b.Plug(New())
		b.Plug(plug.HandleConnFunc(func(conn plug.Conn) {
			assert.Equal(t, conn.Request.Method, c.expected)
		}))
		b.BuildHTTPHandler().ServeHTTP(nil, req)
	}
}
