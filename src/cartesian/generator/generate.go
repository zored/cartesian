package generator

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
)

func Generate(ctx configs.Context, g Generator) (r abstract.ReflectValues, err error) {
	s, err := g.State(ctx)
	if err != nil {
		return nil, err
	}
	for !g.Done(s) {
		next, err := g.Next(s)
		if err != nil {
			return nil, err
		}
		r = append(r, next)
	}
	return r, err
}
