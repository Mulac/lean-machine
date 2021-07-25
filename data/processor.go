package data

type Processor func()

func (p Processor) Process() {
	p()
}

/*
type Processor interface {
	GetName() string
	Process()
}

func NewProcessor(name string, callback func()) Processor {
	return &processor{Name: name, Callback: callback}
}

type processor struct {
	Name string
	Callback func()
}

func (p *processor) GetName() string {
	return p.Name
}

func (p *processor) Process() {
	p.Callback()
}
*/