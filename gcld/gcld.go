package gcld

type Bot struct {
}

func NewBot() Bot {
	return Bot{}
}

func (receiver Bot) Handle(msg interface{}) {
	switch msg.(type) {
	case sendData:

	case recData:

	default:
		panic("error.")
	}
}
