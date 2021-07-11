package configs

type (
	// EntityTemplate must be a pointer on nil entity.
	EntityTemplate interface{}
	Config         struct {
		// EntityTemplate is a pointer to nil value of entity you want to generate.
		EntityTemplate EntityTemplate

		// Fields define how each field must be filled.
		Fields Fields

		// PutIO stores IO after being generated.
		PutIO UpdateIO
	}
	Configs []*Config
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
