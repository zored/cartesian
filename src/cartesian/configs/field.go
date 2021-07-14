package configs

import "github.com/zored/cartesian/src/cartesian/abstract"

type (
	Fields interface {
		GetIOs() IOs
		CreateEntityValues(Context) (r abstract.EntityValues, err error)
		Len() int
		Index(i int) Field
	}
	Field interface {
		CreateValues(io Context) (r abstract.Values, err error)
		GetName() string
		GetParentValue(abstract.Value) abstract.Value
	}
)
