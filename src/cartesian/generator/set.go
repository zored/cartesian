package generator

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"reflect"
)

type Set struct {
	vs   abstract.Values
	i    int
	done bool
}

func NewSet(vs ...abstract.Value) *Set {
	return &Set{vs: abstract.ToValues(vs)}
}

func (s *Set) Done() bool {
	return s.done
}

func (s *Set) Next() (v reflect.Value) {
	v, s.done = s.vs.Index(s.i)
	s.i++
	return v
}
