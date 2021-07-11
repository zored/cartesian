package config

import "github.com/zored/cartesian/src/cartesian/abstract"

type (
	// EntityTemplate must be a pointer on nil entity.
	EntityTemplate interface{}
	Config         struct {
		// EntityTemplate is a pointer to nil value of entity you want to generate.
		EntityTemplate EntityTemplate

		// Fields define how each field must be filled.
		Fields Fields
		PutIO  UpdateIO
	}
	Configs []*Config
	Fields  interface {
		GetIOs() IOs
		CreateEntityValues(*Context) abstract.EntityValues
		Len() int
		Index(i int) Field
	}
	Field interface {
		CreateValues(io *Context) abstract.Values
		GetName() string
		GetParentValue(abstract.Value) abstract.Value
	}
)

func (c *Config) Flatten(includeSelf bool) IOs {
	if !includeSelf {
		return c.Fields.GetIOs()
	}
	r := IOs{}
	io := NewSimpleIO(c)
	r = c.Fields.GetIOs().WithParentIO(io)
	r = append(r, io)
	return r
}
