package cartesian

import "github.com/zored/cartesian/src/cartesian/tag"

// EntityTemplate must be a pointer on nil entity.
type EntityTemplate interface{}

type Config struct {
	// EntityTemplate is a pointer to nil value of entity you want to generate.
	EntityTemplate EntityTemplate

	// Tags define tags `cartesian:"name"` behaviour.
	Tags tag.Tags
}
