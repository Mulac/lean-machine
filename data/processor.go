package data

type Processor interface {
	GetName() string
	Process(m Message) Message
}

type ProcessorBuilder struct{}
