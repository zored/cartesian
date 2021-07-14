package configs

import "reflect"

type (
	// EntityTemplate must be a pointer on nil entity.
	EntityTemplate interface{}
	Config         struct {
		// Name is optional name
		Name string

		// EntityTemplate is a pointer to nil value of entity you want to generate.
		EntityTemplate EntityTemplate `json:"-"`

		// Fields define how each field must be filled.
		Fields Fields

		// PutEntities stores IO after being generated.
		PutEntities PutEntities `json:"-"`
	}
	Configs []*Config
)

func (c *Config) FillName() {
	if c.Name != "" {
		return
	}
	c.Name = "Config for " + reflect.TypeOf(c.EntityTemplate).Elem().Name()
}
func (c *Config) Flatten(includeSelf bool) IOs {
	if !includeSelf {
		return c.Fields.GetIOs()
	}
	io := NewSimpleIO(c)
	return append(c.Fields.GetIOs().WithParentIO(io), io)
}
