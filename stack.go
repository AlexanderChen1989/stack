package stack

import "net/http"

type stack struct {
	frames []FrameFunc
	nexts  []http.HandlerFunc
}

func newStack(frames ...FrameFunc) *stack {
	pipe := &stack{frames: frames}
	pipe.build()
	return pipe
}

// build build nexts for frames
func (s *stack) build() {
	if len(s.frames) == 0 {
		panic("No frame")
	}

	nexts := []http.HandlerFunc{}
	last := len(s.frames) - 1

	for i := 0; i <= last; i++ {
		// last frame has no next
		if i >= last {
			nexts = append(
				nexts,
				func(w http.ResponseWriter, r *http.Request) {},
			)
			break
		}

		nextIndex := i + 1
		nexts = append(
			nexts,
			func(w http.ResponseWriter, r *http.Request) {
				s.frames[nextIndex](w, r, s.nexts[nextIndex])
			},
		)
	}

	s.nexts = nexts
}

func (s *stack) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.frames[0](w, r, s.nexts[0])
}
