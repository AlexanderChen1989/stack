package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRequestID(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("", "", nil)
	next := func(w http.ResponseWriter, r *http.Request) {}

	fakeID := randString(32)
	r.Header.Add(DefaultHTTPHeader, fakeID)
	New("")(w, r, next)
	assert.Equal(t, w.Header().Get(DefaultHTTPHeader), fakeID)

	r.Header.Set(DefaultHTTPHeader, "")
	New("")(w, r, next)
	assert.Len(t, w.Header().Get(DefaultHTTPHeader), 32)
}
