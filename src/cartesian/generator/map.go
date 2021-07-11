package generator

import (
	"github.com/zored/cartesian/src/cartesian/generator/state"
	"reflect"
)

type (
	mapper struct {
		Generator
		mapValue mapValue
	}
	mapValue func(reflect.Value) reflect.Value
)

func NewMap(g Generator, f mapValue) Generator {
	return &mapper{
		Generator: g,
		mapValue:  f,
	}
}

func (f *mapper) Next(st state.State) reflect.Value {
	return f.mapValue(f.Generator.Next(st))
}
