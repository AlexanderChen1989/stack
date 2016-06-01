package gplog

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/AlexanderChen1989/plug"
	"github.com/stretchr/testify/assert"
)

func TestLogPlug(t *testing.T) {
	b := plug.NewBuilder()
	b.Plug(New())
	b.Plug(plug.HandleConnFunc(func(conn plug.Conn) {
		logger := Logger(conn)
		assert.NotNil(t, logger)
		logger.Alert("nice", "Hello, world!")
		logger.Info("nice", "Hello, world!")
	}))
	b.Build().HandleConn(plug.Conn{
		Context: context.Background(),
	})
}
