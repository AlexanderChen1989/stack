package gplog

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/AlexanderChen1989/plug"
)

func TestTrance(t *testing.T) {
	b := plug.NewBuilder()
	b.Plug(New())
	b.Plug(NewTrace())
	b.Build().HandleConn(plug.Conn{
		Context: context.Background(),
	})
}
