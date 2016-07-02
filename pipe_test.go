package plug

import (
	"fmt"
	"net/http"
	"testing"
)

type M struct {
	name string
}

func (m *M) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(nil, nil)
}

func mgen(name string) PlugFunc {
	if name != "" {
		return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			fmt.Println(name)
			next(nil, nil)
		}
	}

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		next(nil, nil)
	}
}

func BenchmarkPlugFunc(b *testing.B) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		builder.PlugFunc(mgen(""))
	}

	h := builder.Build()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(nil, nil)
	}
}

func BenchmarkPlug(b *testing.B) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		builder.Plug(&M{})
	}

	h := builder.Build()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(nil, nil)
	}
}
