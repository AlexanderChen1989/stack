package stack

import "net/http"

// NewBuilder create new Builder
func NewBuilder() *Builder {
	return &Builder{}
}

// Builder is builder for Stack
type Builder struct {
	frames []FrameFunc
}

// Push push Frame to Stack
func (b *Builder) Push(frames ...Frame) {
	for _, f := range frames {
		b.frames = append(b.frames, f.ServeHTTP)
	}
}

// PushFunc push FrameFunc to Stack
func (b *Builder) PushFunc(frames ...FrameFunc) {
	b.frames = append(b.frames, frames...)
}

// Build build Stack
func (b *Builder) Build() http.Handler {
	return newStack(b.frames...)
}
