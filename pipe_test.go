package plug

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

type P struct {
	name string
}

func (m *P) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(nil, nil)
}

func mgen() PlugFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		next(nil, nil)
	}
}

func mgen2(t *testing.T, id int) PlugFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		for i := 0; i <= id; i++ {
			v := r.Context().Value(i)
			if v != nil && !reflect.DeepEqual(v, i) {
				t.Errorf("PID: %d -> ID: %d\n", v, i)
				return
			}
		}

		ctx := context.WithValue(r.Context(), id, id)
		next(nil, r.WithContext(ctx))
	}
}

func TestPlugFunc(t *testing.T) {
	builder := NewBuilder()

	for i := 0; i < 1000; i++ {
		builder.PlugFunc(mgen2(t, i))
	}

	r, _ := http.NewRequest("", "", nil)

	builder.Build().ServeHTTP(nil, r)
}

// BenchmarkPlugFunc for 10 plugs
func BenchmarkPlugFunc(b *testing.B) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		builder.PlugFunc(mgen())
	}

	h := builder.Build()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(nil, nil)
	}
}

// BenchmarkPlug for 10 plugs
func BenchmarkPlug(b *testing.B) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		builder.Plug(&P{})
	}

	h := builder.Build()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(nil, nil)
	}
}
