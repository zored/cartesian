package abstract

import (
	"reflect"
)

type ValuePtr *interface{}
type Value interface{}
type Values []interface{}
type ValuesV []Value

func (v Values) Index(i int) (r reflect.Value, last bool) {
	return reflect.ValueOf(v[i]), len(v)-1 == i
}

func ToValues(v []Value) (r Values) {
	for _, o := range v {
		r = append(r, o)
	}
	return
}
