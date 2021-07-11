package configs

import "github.com/zored/cartesian/src/cartesian/abstract"

type (
	// IO stores input (config), output (entities) and context (parent IO) of generation.
	IO interface {
		GetInput() *Config
		GetOutput() abstract.Entities
		SetOutput(abstract.Entities)
		SetParentIO(IO)
	}
	SimpleIO struct {
		Input    *Config
		Output   abstract.Entities
		ParentIO IO
	}
	UpdateIO func(io IO)
)

func NewSimpleIO(input *Config) *SimpleIO {
	return &SimpleIO{Input: input}
}

func (s *SimpleIO) GetInput() *Config {
	return s.Input
}

func (s *SimpleIO) GetOutput() abstract.Entities {
	return s.Output
}

func (s *SimpleIO) SetOutput(o abstract.Entities) {
	s.Output = o
}

func (s *SimpleIO) SetParentIO(io IO) {
	s.ParentIO = io
}
