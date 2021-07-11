package fields

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/generator"
	"reflect"
)

type Field struct {
	Name         string
	Generator    generator.Generator
	MapParent    MapParent
	InitialValue abstract.Value
}
type MapParent func(parent abstract.Value) (fieldValue abstract.Value)

func NewGenerated(name string, generator generator.Generator) *Field {
	return &Field{Name: name, Generator: generator}
}

func NewFromParent(name string, initial abstract.Value, m MapParent) *Field {
	return &Field{Name: name, MapParent: m, InitialValue: initial}
}

func (f *Field) GetName() string {
	return f.Name
}

func (f *Field) CreateValues(ctx *configs.Context) (r abstract.Values) {
	if f.Generator == nil {
		r = append(r, reflect.ValueOf(f.InitialValue))
		return r
	}
	if f.InitialValue != nil {
		panic("you must fill either field Generator or InitialValue but not both")
	}
	r = generator.Generate(ctx, f.Generator).ToValues()
	return r
}

func (f *Field) GetParentValue(parent abstract.Value) abstract.Value {
	if f.MapParent == nil {
		return parent
	}
	return f.MapParent(parent)
}
