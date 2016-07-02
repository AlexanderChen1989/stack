package plug

import "net/http"

func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct {
	plugs []func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))
}

func (b *Builder) Plug(plugs ...func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))) {
	b.plugs = append(b.plugs, plugs...)
}

func (b *Builder) Build() http.Handler {
	return newPipe(b.plugs...)
}
