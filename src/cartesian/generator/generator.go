package generator

import (
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/generator/state"
	"reflect"
)

type (
	Generator interface {
		State(*configs.Context) state.State
		Next(st state.State) reflect.Value
		Done(st state.State) bool
		GetIOs() configs.IOs
	}
)
