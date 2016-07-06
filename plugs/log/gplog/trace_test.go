package gplog

import (
	"net/http"
	"testing"

	"github.com/AlexanderChen1989/plug"
)

func TestTrance(t *testing.T) {
	b := plug.NewBuilder()
	b.PlugFunc(New())
	b.PlugFunc(Trace)
	r, _ := http.NewRequest("GET", "/", nil)
	b.Build().ServeHTTP(nil, r)
}
