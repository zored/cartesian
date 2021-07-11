package state

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"reflect"
)

type values struct {
	values abstract.Values
	i      int
}

func NewValues(v abstract.Values) *values {
	return &values{values: v}
}

func AsValues(s State) *values {
	return s.(*values)
}

func (s *values) Done() bool {
	return s.i == len(s.values)
}

func (s *values) Next() reflect.Value {
	i := s.i
	s.i++
	return s.values.ValueOfIndex(i)
}
