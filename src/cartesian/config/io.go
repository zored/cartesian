package config

import "github.com/zored/cartesian/src/cartesian/abstract"

type (
	IO interface {
		GetInput() *Config
		GetOutput() abstract.Entities
		SetOutput(abstract.Entities)
		SetParentIO(IO)
	}
	IOs      []IO
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

func (o IOs) Last() IO {
	i := len(o) - 1
	if i < 0 {
		return nil
	}
	return o[i]
}

func (o IOs) WithParentIO(parent IO) IOs {
	for _, io := range o {
		io.SetParentIO(parent)
	}
	return o
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
