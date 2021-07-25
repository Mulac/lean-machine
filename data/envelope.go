package data

type State string

const (
	STATE_NEW        State = "new"
	STATE_PROCESSING State = "processing"
	STATE_FINISHED   State = "finished"
	STATE_ERROR      State = "error"
)

type Envelope interface {
	GetState() State
	GetMetadata() map[string]interface{}
}
