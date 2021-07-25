package data

type Pipeline interface {
	Execute()
}

type pipeline struct {
	workers	   int
	processors []Processor
}

func (p *pipeline) Execute() {
	for _, processor := range p.processors {
		for i := 0; i < p.workers; i++ {
			go processor.Process()
		}
	}

	select{}
}

type PipelineBuilder struct {
	workers    int
	processors []Processor
}

func NewPipelineBuilder() *PipelineBuilder {
	return new(PipelineBuilder)
}

func (pb *PipelineBuilder) FromJSON() *PipelineBuilder {
	return pb
}

func (pb *PipelineBuilder) AddProcessor(processor Processor) *PipelineBuilder {
	pb.processors = append(pb.processors, processor)
	return pb
}

func (pb *PipelineBuilder) Workers(count int) *PipelineBuilder {
	pb.workers = count
	return pb
}

func (pb *PipelineBuilder) Build() (Pipeline, error) {
	return &pipeline{
		workers:    pb.workers,
		processors: pb.processors,
	}, nil
}
