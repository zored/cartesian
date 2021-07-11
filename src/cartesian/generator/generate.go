package generator

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/config"
)

func Generate(ctx *config.Context, g Generator) (r abstract.ReflectValues) {
	s := g.State(ctx)
	for !g.Done(s) {
		r = append(r, g.Next(s))
	}
	return r
}
