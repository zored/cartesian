package configs

import "github.com/zored/cartesian/src/cartesian/abstract"

type (
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
