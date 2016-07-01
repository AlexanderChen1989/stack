package plug

import "net/http"

type Plug interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type PlugFunc func(http.ResponseWriter, *http.Request, http.HandlerFunc)

func (fn PlugFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fn(w, r, next)
}

type Pipe []Plug

func Next(w http.ResponseWriter, r *http.Request, ps Pipe) {
	if len(ps) <= 0 {
		return
	}

	ps[0].ServeHTTP(
		w, r,
		func(nw http.ResponseWriter, nr *http.Request) {
			Next(nw, nr, ps[1:])
		},
	)
}

func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct {
	pipe Pipe
}

func (b *Builder) Plug(p Plug) {
	b.pipe = append(b.pipe, p)
}

func (b *Builder) PlugFunc(fn func(http.ResponseWriter, *http.Request, http.HandlerFunc)) {
	b.pipe = append(b.pipe, PlugFunc(fn))
}

func (b *Builder) Build() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Next(w, r, b.pipe)
	})
}

func (b *Builder) Pipe() Pipe {
	return b.pipe
}
