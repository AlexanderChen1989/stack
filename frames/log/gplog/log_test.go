package gplog

import (
	"net/http"
	"testing"

	"github.com/AlexanderChen1989/stack"
	"github.com/stretchr/testify/assert"
)

func TestLogPlug(t *testing.T) {
	checkFn := func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		logger := Logger(r)
		assert.NotNil(t, logger)
		logger.Alert("[Nice] ", "Hello, world!")
		logger.Info("[Nice] ", "Hello, world!")
	}
	b := stack.NewBuilder()
	b.PushFunc(New())
	b.PushFunc(checkFn)
	r, _ := http.NewRequest("GET", "/", nil)
	b.Build().ServeHTTP(nil, r)
}
