package fields

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/generator"
	"reflect"
)

type (
	field struct {
		Name      string
		Generator generator.Generator
		Values    abstract.Values
	}
	Fields []*field
)

func NewValued(name string, values ...abstract.Value) *field {
	return &field{Name: name, Values: abstract.ToValues(values)}
}

func NewGenerated(name string, generator generator.Generator) *field {
	return &field{Name: name, Generator: generator}
}

func NewFields(fields ...*field) Fields {
	return fields
}

func (t field) Apply(field reflect.Value) bool {
	reflect.Indirect(field).Set(t.Generator.Next())
	return t.Generator.Done()
}
