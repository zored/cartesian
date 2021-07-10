package fields

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/generator"
)

type (
	Field struct {
		Name      string
		Generator generator.Generator
		Values    abstract.Values
	}
	Fields []*Field
)

func (f Field) MaterializeValues() {
	if f.Values != nil {
		return
	}
	if g := f.Generator; g != nil {
		f.Values = generateValues(g)
	}
}

func (f Fields) MaterializeValues() {
	for _, t := range f {
		t.MaterializeValues()
	}
}

func NewValued(name string, values ...abstract.Value) *Field {
	return &Field{Name: name, Values: abstract.ToValues(values)}
}

func NewGenerated(name string, generator generator.Generator) *Field {
	return &Field{Name: name, Generator: generator}
}

func NewFields(fields ...*Field) Fields {
	return fields
}

func generateValues(g generator.Generator) (r abstract.Values) {
	for !g.Done() {
		r = append(r, g.Next())
	}
	return r
}
