package queue

import (
	"lean-machine/data"
	"time"
)

type naiveQueue struct {
	topics map[data.State]chan data.Message
}

// Put will place a Message on a topic
func (q *naiveQueue) Put(state data.State, messages ...data.Message) {
		for _, msg := range messages {
			q.topics[state] <- msg
		}
}

// Fetch will retrieve a Message from a topic. Block until a message is received.
func (q *naiveQueue) Fetch(state data.State) (data.Message, error) {
	for {
		select {
		case m := <-q.topics[state]:
			return m, nil
		default:
			continue
		}
	}
}

// TryFetch will fetch a Message from a topic. Non-blocking.
func (q *naiveQueue) TryFetch(state data.State) (data.Message, error) {
	select {
	case m := <-q.topics[state]:
		return m, nil
	default:
		return nil, nil
	}
}

//FetchWithTimeout will try to receive a Message from a topic. Blocking and timeout after duration.
func (q *naiveQueue) FetchWithTimeout(state data.State, d time.Duration) (data.Message, error) {
	select {
	case m := <-q.topics[state]:
		return m, nil
	case <-time.After(d):
		return nil, TimeoutError
	}
}
