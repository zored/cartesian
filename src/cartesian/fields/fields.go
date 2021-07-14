package fields

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
)

type Fields []*Field

func (f Fields) Len() int {
	return len(f)
}

func (f Fields) Index(i int) configs.Field {
	if i >= f.Len() || i < 0 {
		return nil
	}
	return f[i]
}
func NewFields(fields ...*Field) Fields {
	return fields
}
func (f Fields) CreateEntityValues(ctx configs.Context) (r abstract.EntityValues, err error) {
	for _, t := range f {
		values, err := t.CreateValues(ctx)
		if err != nil {
			return nil, err
		}
		r = append(r, values)
	}
	return r, err
}

func (f Fields) GetIOs() configs.IOs {
	r := configs.IOs{}
	for _, v := range f {
		r = append(r, v.Generator.GetIOs()...)
	}
	return r
}
