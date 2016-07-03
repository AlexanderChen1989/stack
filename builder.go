package plug

import "net/http"

// NewBuilder create new Builder
func NewBuilder() *Builder {
	return &Builder{}
}

// Builder is builder for Pipe
type Builder struct {
	plugs []PlugFunc
}

// Plug plug Plug to Pipe
func (b *Builder) Plug(plugs ...Plug) {
	for _, p := range plugs {
		b.plugs = append(b.plugs, p.ServeHTTP)
	}
}

// PlugFunc plug PlugFunc to Pipe
func (b *Builder) PlugFunc(plugs ...PlugFunc) {
	b.plugs = append(b.plugs, plugs...)
}

// Build build Pipe
func (b *Builder) Build() http.Handler {
	return newPipe(b.plugs...)
}
