package plug

import (
	"net/http"

	"golang.org/x/net/context"
)

type Builder struct {
	pluggers []Plugger
}

func NewBuilder(ps ...Plugger) *Builder {
	base := []Plugger{emptyPlugger{}}
	return &Builder{
		pluggers: append(base, ps...),
	}
}

func (builder *Builder) Plug(plug Plugger) {
	builder.pluggers = append(builder.pluggers, plug)
}

func (builder *Builder) Build() Plugger {
	p := builder.pluggers[0]

	for i := len(builder.pluggers) - 1; i >= 1; i-- {
		p = builder.pluggers[i].Plug(p)
	}

	return p
}

func (builder *Builder) BuildHTTPHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn := Conn{
			Context:        context.Background(),
			Request:        r,
			ResponseWriter: w,
		}

		builder.Build().Handle(conn)
	})
}
