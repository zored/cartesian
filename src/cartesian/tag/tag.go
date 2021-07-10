package tag

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/generator"
	"reflect"
)

type (
	Tag struct {
		Name      string
		Generator generator.Generator
		Values    abstract.Values
	}
	Tags []*Tag
)

func NewValued(name string, values ...abstract.Value) *Tag {
	return &Tag{Name: name, Values: abstract.ToValues(values)}
}

func NewGenerated(name string, generator generator.Generator) *Tag {
	return &Tag{Name: name, Generator: generator}
}

func NewTags(tags ...*Tag) Tags {
	return tags
}

func (t Tag) Apply(field reflect.Value) bool {
	reflect.Indirect(field).Set(t.Generator.Next())
	return t.Generator.Done()
}
