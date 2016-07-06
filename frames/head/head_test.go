package head

import (
	"net/http"
	"testing"

	"github.com/AlexanderChen1989/stack"
	"github.com/stretchr/testify/assert"
)

func TestHeadFrame(t *testing.T) {
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

		checkFn := func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			assert.Equal(t, r.Method, c.expected)
		}

		b := stack.NewBuilder()
		b.PushFunc(FrameFunc)
		b.PushFunc(checkFn)

		b.Build().ServeHTTP(nil, req)
	}
}
