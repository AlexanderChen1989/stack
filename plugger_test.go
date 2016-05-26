package plug

import (
	"fmt"
	"testing"
)

type helloPlugger struct {
	next Plugger
}

//Plug implements PluggerPlug
func (hello *helloPlugger) Plug(p Plugger) Plugger {
	hello.next = p
	return hello
}

// Handle implements Plugger.Handle
func (hello *helloPlugger) Handle(conn Conn) {
	fmt.Println("Hello")
	hello.next.Handle(conn)
}

func TestBuilder(t *testing.T) {
	b := NewBuilder()

	b.Plug(&helloPlugger{})

	b.Build().Handle(Conn{})
}
