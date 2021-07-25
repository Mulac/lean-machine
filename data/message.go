package data

type Message interface {
	Envelope() Envelope
	Payload() Payload
}
