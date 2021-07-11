package generator

import (
	"github.com/zored/cartesian/src/cartesian/config"
	"github.com/zored/cartesian/src/cartesian/generator/state"
	"reflect"
)

type (
	groupper struct {
		Generator
	}
	doneState struct {
		done bool
		ctx  *config.Context
	}
)

func NewGroup(g Generator) Generator {
	return &groupper{
		Generator: g,
	}
}

func (g *groupper) State(ctx *config.Context) state.State {
	return &doneState{ctx: ctx}
}

func (g *groupper) Next(st state.State) reflect.Value {
	s := st.(*doneState)
	s.done = true
	return Generate(s.ctx, g.Generator).ToValueListReflection()
}

func (g *groupper) Done(st state.State) bool {
	return st.(*doneState).done
}
