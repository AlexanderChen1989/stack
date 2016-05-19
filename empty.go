package plug

import "fmt"

type emptyPlugger struct{}

func (empty emptyPlugger) Plug(Plugger) Plugger {
	return empty
}

func (_ emptyPlugger) Handle(Conn) {
	fmt.Println("Empty")
}
