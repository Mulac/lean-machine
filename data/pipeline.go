package pipe

type Pipeline interface {
	Execute()
}

type Processor interface {
	GetName() string
	Process(m Message) Message
}
