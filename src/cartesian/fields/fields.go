package fields

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/config"
)

type Fields []*Field

func (f Fields) Len() int {
	return len(f)
}

func (f Fields) Index(i int) config.Field {
	return f[i]
}
func NewFields(fields ...*Field) Fields {
	return fields
}
func (f Fields) CreateEntityValues(ctx *config.Context) (r abstract.EntityValues) {
	for _, t := range f {
		r = append(r, t.CreateValues(ctx))
	}
	return r
}

func (f Fields) GetIOs() config.IOs {
	r := config.IOs{}
	for _, v := range f {
		r = append(r, v.Generator.GetIOs()...)
	}
	return r
}
