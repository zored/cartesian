package configs

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"reflect"
)

type (
	// IO stores input (config), output (entities) and context (parent IO) of generation.
	IO interface {
		GetInput() *Config
		SetParentIO(IO)
	}
	SimpleIO struct {
		EntityTemplateName string
		Input              *Config
		Output             abstract.Entities
		ParentIO           IO
	}
)

func NewSimpleIO(input *Config) *SimpleIO {
	return &SimpleIO{
		EntityTemplateName: reflect.TypeOf(input.EntityTemplate).Elem().Name(),
		Input:              input,
	}
}

func (s *SimpleIO) GetInput() *Config {
	return s.Input
}

func (s *SimpleIO) GetOutput() abstract.Entities {
	return s.Output
}

func (s *SimpleIO) SetParentIO(io IO) {
	s.ParentIO = io
}
