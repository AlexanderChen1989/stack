package rmux

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexanderChen1989/plug"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestMux(t *testing.T) {
	router := mux.NewRouter()

	genHandle := func(path string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, path, r.URL.Path)
		}
	}

	paths := []string{
		"/hello",
		"/world",
		"/hello/world",
	}
	for _, path := range paths {
		router.HandleFunc(path, genHandle(path))
	}

	b := plug.NewBuilder()
	b.PlugFunc(New(router))

	server := httptest.NewServer(b.Build())

	defer server.Close()

	for _, path := range paths {
		http.Get(server.URL + path)
	}
}
