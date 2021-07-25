package queue

import (
	"errors"
	"lean-machine/data"
	"time"
)

var (
	TimeoutError = errors.New("timeout fetching message from topic")
)


type Queue interface {
	Fetch(state data.State) (data.Message, error)
	TryFetch(state data.State) (data.Message, error)
	FetchWithTimeout(state data.State, d time.Duration) (data.Message, error)
	Put(state data.State, messages ...data.Message)
}