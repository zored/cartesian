package configs

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
)

type (
	Context struct {
		ContextLocal `json:"-"`
		*ContextGlobal
	}
	ContextLocal struct {
		LocalResult LocalResult
	}
	ContextGlobal struct {
		Factories  *TemplateFactories
		AllResults *ResultVisitor
	}
	PutEntities func(ctx Context, entities abstract.Entities)
)

func NewContext() Context {
	visitor, localResult := NewResultVisitor()
	return Context{
		ContextLocal: ContextLocal{
			LocalResult: localResult,
		},
		ContextGlobal: &ContextGlobal{
			Factories:  &TemplateFactories{},
			AllResults: visitor,
		},
	}
}

func (c Context) EachFactory(ok func(*TemplateFactory) (stop bool)) {
	for _, f := range *c.Factories {
		if ok(f) {
			return
		}
	}
}

func (c Context) WithFactories(factories TemplateFactories) Context {
	*c.Factories = factories
	return c
}

func (c Context) WithConfig(config *Config) Context {
	c.AllResults.SetConfig(config, &c.LocalResult)
	return c
}

func (c Context) WithEntity(v abstract.Entity) Context {
	c.AllResults.SetEntity(v, &c.LocalResult)
	return c
}

func (c Context) WithField(v Field) Context {
	c.AllResults.SetField(v, &c.LocalResult)
	return c
}

func (c Context) WithFieldValuePointer(v interface{}) Context {
	c.AllResults.SetFieldValuePointer(v, &c.LocalResult)
	return c
}
