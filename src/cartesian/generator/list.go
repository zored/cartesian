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

func NewList(l ...abstract.Value) *list {
	return &list{l: abstract.ToValues(l)}
}

func (s *list) State(*configs.Context) state.State {
	return state.NewValues(s.l)
}

func (s *list) Done(st state.State) bool {
	return state.AsValues(st).Done()
}

func (s *list) Next(st state.State) reflect.Value {
	return state.AsValues(st).Next()
}

func (s *list) GetIOs() (r configs.IOs) {
	return r
}
