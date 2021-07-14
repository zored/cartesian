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

func (s *entity) State(ctx configs.Context) (r state.State, err error) {
	var factory *configs.TemplateFactory
	ctx.EachFactory(func(v *configs.TemplateFactory) bool {
		if o, ok := v.IO.(*io); ok && o.generatorId == s.generatorId {
			factory = v
			return true
		}
		return false
	})
	if factory == nil {
		panic("no factory with IO found for generator (it must be generated in previous config iteration)")
	}
	entities, err := factory.Create(ctx)
	if err != nil {
		return nil, err
	}
	values := entities.AsValues()
	if s.c.List {
		return state.NewValues(abstract.Values{values}), err
	}

	return state.NewValues(values), err
}

func (s *entity) Done(st state.State) bool {
	return state.AsValues(st).Done()
}

func (s *entity) Next(st state.State) (v reflect.Value, err error) {
	return state.AsValues(st).Next(), nil
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
