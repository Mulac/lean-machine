package lean_machine

import (
	"lean-machine/data"
	"lean-machine/queue"
)

func DoSomethingCoolProcessor() {
	for {
		msg, err := queue.GetQueue().Fetch("mytopic")
		if err != nil {
			// TODO
		}

		queue.GetQueue().Put("nextState", msg)
	}
}

func main() {
	machine, _ := data.NewPipelineBuilder().AddProcessor(DoSomethingCoolProcessor).Workers(10).Build()
	machine.Execute()
}
