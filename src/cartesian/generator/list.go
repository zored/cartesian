package generator

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/config"
	"github.com/zored/cartesian/src/cartesian/generator/state"
	"reflect"
)

type list struct {
	l abstract.Values
}

func NewList(l ...abstract.Value) *list {
	return &list{l: abstract.ToValues(l)}
}

func (s *list) State(*config.Context) state.State {
	return state.NewValues(s.l)
}

func (s *list) Done(st state.State) bool {
	return state.AsValues(st).Done()
}

func (s *list) Next(st state.State) reflect.Value {
	return state.AsValues(st).Next()
}

func (s *list) GetIOs() (r config.IOs) {
	return r
}
