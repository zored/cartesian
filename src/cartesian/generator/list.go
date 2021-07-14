package generator

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/generator/state"
	"reflect"
)

type list struct {
	l abstract.Values
}

func NewList(l ...interface{}) *list {
	return &list{l: l}
}

func (s *list) State(configs.Context) (state.State, error) {
	return state.NewValues(s.l), nil
}

func (s *list) Done(st state.State) bool {
	return state.AsValues(st).Done()
}

func (s *list) Next(st state.State) (reflect.Value, error) {
	return state.AsValues(st).Next(), nil
}

func (s *list) GetIOs() (r configs.IOs) {
	return r
}
