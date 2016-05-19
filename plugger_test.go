package plug

import (
	"fmt"
	"testing"
)

type helloPlugger struct {
	next Plugger
}

func (hello *helloPlugger) Plug(p Plugger) Plugger {
	hello.next = p
	return hello
}

func (p *helloPlugger) Handle(conn Conn) {
	fmt.Println("Hello")
	p.next.Handle(conn)
}

func TestBuilder(t *testing.T) {
	b := NewBuilder()

	b.Plug(&helloPlugger{})

	b.Build().Handle(Conn{})
}
