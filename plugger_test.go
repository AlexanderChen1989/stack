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

// HandleConn implements Plugger.HandleConn
func (hello *helloPlugger) HandleConn(conn Conn) {
	fmt.Println("Hello")
	hello.next.HandleConn(conn)
}

func TestBuilder(t *testing.T) {
	b := NewBuilder()

	b.Plug(&helloPlugger{})

	b.Build().HandleConn(Conn{})
}
