package generator

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/generator/state"
	"reflect"
)

var lastEntityGeneratorId = 0

type (
	entity struct {
		c           *EntityConfig
		generatorId int
	}
	io struct {
		*configs.SimpleIO
		generatorId int
	}
	EntityConfig struct {
		GeneratorConfig *configs.Config
		List            bool
	}
)

func NewEntityList(c *configs.Config) *entity {
	return NewEntity(&EntityConfig{
		GeneratorConfig: c,
		List:            true,
	})
}

func NewEntitySingle(c *configs.Config) *entity {
	return NewEntity(&EntityConfig{
		GeneratorConfig: c,
	})
}

func NewEntity(c *EntityConfig) *entity {
	lastEntityGeneratorId++
	return &entity{
		c:           c,
		generatorId: lastEntityGeneratorId,
	}
}

func (s *entity) State(ctx *configs.Context) (r state.State) {
	var generatorIO configs.IO
	ctx.EachCompleteIO(func(v configs.IO) bool {
		if o, ok := v.(*io); ok && o.generatorId == s.generatorId {
			generatorIO = o
			return true
		}
		return false
	})
	if generatorIO == nil {
		panic("no IO found for generator (it must be generated in previous config iteration)")
	}
	values := generatorIO.GetOutput().AsValues()
	if s.c.List {
		return state.NewValues(abstract.Values{values})
	}
	return state.NewValues(values)
}

func (s *entity) Done(st state.State) bool {
	return state.AsValues(st).Done()
}

func (s *entity) Next(st state.State) (v reflect.Value) {
	return state.AsValues(st).Next()
}

func (s *entity) GetIOs() configs.IOs {
	c := s.c.GeneratorConfig
	r := configs.IOs{}
	r = append(r, c.Flatten(false)...)
	r = append(r, &io{
		SimpleIO:    configs.NewSimpleIO(c),
		generatorId: s.generatorId,
	})
	return r
}
