package stack

import "net/http"

type stack struct {
	frames []FrameFunc
	nexts  []http.HandlerFunc
}

func newStack(frames ...FrameFunc) *stack {
	pipe := &stack{frames: frames}
	pipe.buildNexts()
	return pipe
}

func (s *stack) buildNexts() {
	if len(s.frames) == 0 {
		panic("No plug")
	}

	frames := s.frames
	nexts := []http.HandlerFunc{
		func(w http.ResponseWriter, r *http.Request) {},
	}

	for i := len(frames) - 2; i >= 0; i-- {
		index := i + 1
		next := func(w http.ResponseWriter, r *http.Request) {
			s.frames[index](w, r, s.nexts[index])
		}
		nexts = append([]http.HandlerFunc{next}, nexts...)
	}

	s.nexts = nexts
}

func (s *stack) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.frames[0](w, r, s.nexts[0])
}
