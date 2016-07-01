package plug

import (
	"fmt"
	"net/http"
	"testing"
)

func genPlug(name string) Plug {
	return PlugFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Println(name)
		next(w, r)
	})
}

func BenchmarkPlug(b *testing.B) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		builder.PlugFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			next(w, r)
		})
	}

	handler := builder.Build()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(nil, nil)
	}
}

func BenchmarkNoPlug(b *testing.B) {
	fn := func(w http.ResponseWriter, r *http.Request) {}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			fn(nil, nil)
		}
	}
}

func BenchmarkCreateFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			_ = func(w http.ResponseWriter, r *http.Request) {}
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			_ = p[1:]
		}
	}
}

func xTestBuilder(t *testing.T) {
	b := NewBuilder()
	b.Plug(genPlug("P1"))
	b.Plug(genPlug("P2"))
	b.Plug(genPlug("P3"))
	b.Plug(genPlug("P4"))
	b.Plug(genPlug("P5"))
	b.Build().ServeHTTP(nil, nil)
}
