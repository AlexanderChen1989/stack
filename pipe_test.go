package plug

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

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

func TestNext(t *testing.T) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		name := strconv.Itoa(i)
		builder.Plug(mgen(name))
	}

	builder.Build().ServeHTTP(nil, nil)
}

func BenchmarkNext(b *testing.B) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		builder.Plug(mgen(""))
	}

	h := builder.Build()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(nil, nil)
	}
}
