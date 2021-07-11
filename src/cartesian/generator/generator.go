package generator

import (
	"github.com/zored/cartesian/src/cartesian/config"
	"github.com/zored/cartesian/src/cartesian/generator/state"
	"reflect"
)

type (
	Generator interface {
		State(*config.Context) state.State
		Next(st state.State) reflect.Value
		Done(st state.State) bool
		GetIOs() config.IOs
	}
)
