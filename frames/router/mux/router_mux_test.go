package mux

import (
	"net/http"
	"net/http/httptest"
	"testing"

	rmux "github.com/gorilla/mux"

	"github.com/AlexanderChen1989/plug"
	"github.com/AlexanderChen1989/plug/plugs/router"
	"github.com/stretchr/testify/assert"
)

func TestMux(t *testing.T) {
	r := rmux.NewRouter()

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
		r.HandleFunc(path, genHandle(path))
	}

	b := plug.NewBuilder()
	b.PlugFunc(router.New(r))

	server := httptest.NewServer(b.Build())

	defer server.Close()

	for _, path := range paths {
		http.Get(server.URL + path)
	}
}
