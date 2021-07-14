package generator

import (
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/generator/state"
	"reflect"
)

type (
	groupper struct {
		Generator
	}
	doneState struct {
		done bool
		ctx  configs.Context
	}
)

func NewGroup(g Generator) Generator {
	return &groupper{
		Generator: g,
	}
}

func (g *groupper) State(ctx configs.Context) (state.State, error) {
	return &doneState{ctx: ctx}, nil
}

func (g *groupper) Next(st state.State) (reflect.Value, error) {
	s := st.(*doneState)
	s.done = true
	entities, err := Generate(s.ctx, g.Generator)
	if err != nil {
		return reflect.Value{}, err
	}
	return entities.ToValueListReflection(), nil
}

func (g *groupper) Done(st state.State) bool {
	return st.(*doneState).done
}
