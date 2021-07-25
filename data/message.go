package pipe

type Message interface {
	Envelope() Envelope
	Payload() Payload
}

type Envelope interface {
	GetState() State
	GetMetadata() map[string]interface{}
}

type Payload interface {
	Type() PayloadType
	Value() interface{}
}

type State string

const (
	STATE_NEW        State = "new"
	STATE_PROCESSING State = "processing"
	STATE_FINISHED   State = "finished"
	STATE_ERROR      State = "error"
)

type PayloadType string
