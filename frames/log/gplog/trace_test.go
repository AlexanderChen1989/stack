package gplog

import (
	"net/http"
	"testing"

	"github.com/AlexanderChen1989/stack"
)

func TestTrance(t *testing.T) {
	b := stack.NewBuilder()
	b.PushFunc(New())
	b.PushFunc(Trace)
	r, _ := http.NewRequest("GET", "/", nil)
	b.Build().ServeHTTP(nil, r)
}
