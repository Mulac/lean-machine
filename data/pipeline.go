package data

type Pipeline interface {
	Execute()
}

type PipelineBuilder struct{}

func (pb *PipelineBuilder) FromJSON() PipelineBuilder {
	return *pb
}

func (pb *PipelineBuilder) Build() (Pipeline, error) {
	return nil, nil
}
