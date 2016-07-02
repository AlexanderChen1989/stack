package plug

import "net/http"

type Builder struct {
	plugs []func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))
}

func (b *Builder) Plug(plugs ...func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))) {
	b.plugs = plugs
}

func (b *Builder) Build() http.Handler {
	return newPipe(b.plugs...)
}
