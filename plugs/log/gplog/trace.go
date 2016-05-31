package gplog

import (
	"fmt"

	"github.com/AlexanderChen1989/plug"
)

type tracePlug struct {
	next plug.Plugger
}

// NewTrace create trace plugger to trace request consumed time
func NewTrace() plug.Plugger {
	return &tracePlug{}
}

func (tr *tracePlug) Plug(next plug.Plugger) plug.Plugger {
	tr.next = next
	return tr
}

func (tr *tracePlug) HandleConn(conn plug.Conn) {
	logger := Logger(conn)
	if logger == nil {
		fmt.Println("Please add log plug first")
	} else {
		defer logger.Trace("[Request]").End()
	}

	tr.next.HandleConn(conn)
}
