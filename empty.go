package plug

type emptyPlugger struct{}

func NewEmpty() Plugger {
	return emptyPlugger{}
}

func (empty emptyPlugger) Plug(Plugger) Plugger {
	return empty
}

func (_ emptyPlugger) Handle(Conn) {
}
