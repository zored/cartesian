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
	io := NewSimpleIO(c)
	return append(c.Fields.GetIOs().WithParentIO(io), io)
}
