package plug

import "net/http"

type pipe struct {
	plugs []func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))
	nexts []func(http.ResponseWriter, *http.Request)
}

func newPipe(plugs ...func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))) *pipe {
	return &pipe{plugs: plugs}
}

func emptyPlug(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request)) {}

func (p *pipe) build() {
	if len(p.plugs) == 0 {
		panic("No plug")
	}

	plugs := p.plugs

	nexts := []func(http.ResponseWriter, *http.Request){
		func(w http.ResponseWriter, r *http.Request) {},
	}

	for i := len(plugs) - 2; i >= 0; i-- {
		rest := plugs[i+1:]
		next := func(w http.ResponseWriter, r *http.Request) {
			p.next(w, r, rest)
		}
		head := []func(w http.ResponseWriter, r *http.Request){next}
		nexts = append(head, nexts...)
	}

	p.nexts = nexts
}

func (p *pipe) next(w http.ResponseWriter, r *http.Request,
	plugs []func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))) {

	nexts := p.nexts

	plugs[0](w, r, nexts[len(nexts)-len(plugs)])
}

func (p *pipe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.next(w, r, p.plugs)
}
