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

func NewMap(g Generator, f func(interface{}) interface{}) Generator {
	return NewMapReflect(g, func(r reflect.Value) reflect.Value {
		return reflect.ValueOf(f(r.Interface()))
	})
}

func NewMapReflect(g Generator, f mapValue) Generator {
	return &mapper{
		Generator: g,
		mapValue:  f,
	}
}

func (f *mapper) Next(st state.State) (reflect.Value, error) {
	next, err := f.Generator.Next(st)
	if err != nil {
		return reflect.Value{}, err
	}
	return f.mapValue(next), nil
}
