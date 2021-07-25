package data

type PayloadType string

type Payload interface {
	Type() PayloadType
	Value() interface{}
}
