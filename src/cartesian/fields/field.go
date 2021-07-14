package fields

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/generator"
	"reflect"
)

type Field struct {
	Name         string
	Generator    generator.Generator `json:"-"`
	MapParent    MapParent           `json:"-"`
	InitialValue abstract.Value      `json:",omitempty"`
}
type MapParent func(parent abstract.Value) (fieldValue abstract.Value)

func NewGenerated(name string, generator generator.Generator) *Field {
	return &Field{Name: name, Generator: generator}
}

func NewList(name string, values ...interface{}) *Field {
	return NewGenerated(name, generator.NewList(values...))
}

func (f *Field) GetName() string {
	return f.Name
}

func (f *Field) CreateValues(ctx configs.Context) (r abstract.Values, err error) {
	if f.Generator == nil {
		r = append(r, reflect.ValueOf(f.InitialValue))
		return r, err
	}
	if f.InitialValue != nil {
		panic("you must fill either field Generator or InitialValue but not both")
	}
	reflectValues, err := generator.Generate(ctx.WithField(f), f.Generator)
	if err != nil {
		return nil, err
	}
	return reflectValues.ToValues(), err
}

func (f *Field) GetParentValue(parent abstract.Value) abstract.Value {
	if f.MapParent == nil {
		return parent
	}
	return f.MapParent(parent)
}
