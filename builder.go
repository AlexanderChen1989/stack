package plug

import "net/http"

func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct {
	plugs []PlugFunc
}

func (b *Builder) Plug(plugs ...Plug) {
	for _, p := range plugs {
		b.plugs = append(b.plugs, p.ServeHTTP)
	}
}

func (b *Builder) PlugFunc(plugs ...PlugFunc) {
	b.plugs = append(b.plugs, plugs...)
}

func (b *Builder) Build() http.Handler {
	return newPipe(b.plugs...)
}
