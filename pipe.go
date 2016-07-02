package plug

import "net/http"

type PlugFunc func(http.ResponseWriter, *http.Request, http.HandlerFunc)

type pipe struct {
	plugs []PlugFunc
	nexts []http.HandlerFunc
}

func newPipe(plugs ...PlugFunc) *pipe {
	pipe := &pipe{plugs: plugs}
	pipe.buildNexts()
	return pipe
}

func (p *pipe) buildNexts() {
	if len(p.plugs) == 0 {
		panic("No plug")
	}

	plugs := p.plugs

	nexts := []http.HandlerFunc{
		func(w http.ResponseWriter, r *http.Request) {},
	}

	for i := len(plugs) - 2; i >= 0; i-- {
		index := i + 1
		next := func(w http.ResponseWriter, r *http.Request) {
			p.plugs[index](w, r, p.nexts[index])
		}

		nexts = append([]http.HandlerFunc{next}, nexts...)
	}

	p.nexts = nexts
}

func (p *pipe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.plugs[0](w, r, p.nexts[0])
}
