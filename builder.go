package plug

import "net/http"

func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct {
	plugs []PlugFunc
}

func (b *Builder) Plug(plugs ...PlugFunc) {
	b.plugs = append(b.plugs, plugs...)
}

func (b *Builder) Build() http.Handler {
	return newPipe(b.plugs...)
}
