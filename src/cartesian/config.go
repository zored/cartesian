package cartesian

import "github.com/zored/cartesian/src/cartesian/fields"

// EntityTemplate must be a pointer on nil entity.
type EntityTemplate interface{}

type Config struct {
	// EntityTemplate is a pointer to nil value of entity you want to generate.
	EntityTemplate EntityTemplate

	// Fields define how each field must be filled.
	Fields fields.Fields
}
