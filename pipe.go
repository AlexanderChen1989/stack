package plug

import "net/http"

type pipe struct {
	plugs []func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))
	nexts []func(http.ResponseWriter, *http.Request)
}

func newPipe(plugs ...func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))) *pipe {
	return &pipe{plugs: plugs}
}

func (p *pipe) build() {
	if len(p.plugs) == 0 {
		panic("No plug")
	}

	plugs := p.plugs

	nexts := []func(http.ResponseWriter, *http.Request){
		func(w http.ResponseWriter, r *http.Request) {},
	}

	for i := len(plugs) - 2; i >= 0; i-- {
		plugFn, index := plugs[i+1], i+1

		next := func(w http.ResponseWriter, r *http.Request) {
			plugFn(w, r, p.nexts[index])
		}

		head := []func(w http.ResponseWriter, r *http.Request){next}

		nexts = append(head, nexts...)
	}

	p.nexts = nexts
}

func (p *pipe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.plugs[0](w, r, p.nexts[0])
}
