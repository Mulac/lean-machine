package queue

import (
	"lean-machine/data"
	"log"
	"os"
	"sync"
)

var (
	queue Queue
	queueOnce sync.Once
)

const EnvQueueImpl = "QUEUE_IMPL"

type QueueType string
const (
	Naive = "naive"
)

func GetQueue() Queue {
	queueOnce.Do(func() {
		impl := os.Getenv(EnvQueueImpl)
		switch impl {
		case Naive, "":
			queue = &naiveQueue{topics: make(map[data.State]chan data.Message)}
		default:
			log.Fatalf("unknown queue implementation: %s", impl)
		}
	})
	return queue
}