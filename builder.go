package plug

import (
	"net/http"

	"golang.org/x/net/context"
)

// Builder build pipeline of Pluggers
type Builder struct {
	pluggers []Plugger
}

// NewBuilder create a new Builder
func NewBuilder(ps ...Plugger) *Builder {
	base := []Plugger{emptyPlugger{}}
	return &Builder{
		pluggers: append(base, ps...),
	}
}

// Plug plug Plugger to Builder
func (builder *Builder) Plug(plug Plugger) {
	builder.pluggers = append(builder.pluggers, plug)
}

// Build build pipeline of Pluggers to a single Plugger
func (builder *Builder) Build() Plugger {
	p := builder.pluggers[0]

	for i := len(builder.pluggers) - 1; i >= 1; i-- {
		p = builder.pluggers[i].Plug(p)
	}

	return p
}

// BuildHTTPHandler build pipeline of Pluggers to a single http.Handler
func (builder *Builder) BuildHTTPHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn := Conn{
			Context:        context.Background(),
			Request:        r,
			ResponseWriter: w,
		}

		builder.Build().HandleConn(conn)
	})
}
