package generator

import (
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/generator/state"
	"reflect"
)

type (
	funk struct {
		f Func
	}
	FuncResult struct {
		Done  bool
		Value interface{}
	}
	FuncResults []*FuncResult
	Func        func() *FuncResult
)

func NewFunc(f Func) *funk {
	return &funk{f: f}
}

func (s *funk) State(configs.Context) (state.State, error) {
	return s.f(), nil
}

func (s *funk) Done(st state.State) bool {
	return st.(*FuncResult).Done
}

func (s *funk) Next(st state.State) (reflect.Value, error) {
	ss := st.(*FuncResult)
	v := ss.Value
	*ss = *s.f()

	return reflect.ValueOf(v), nil
}

func (s *funk) GetIOs() (r configs.IOs) {
	return r
}
