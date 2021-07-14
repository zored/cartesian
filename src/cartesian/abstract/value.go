package abstract

import (
	"reflect"
)

type (
	ValuePtr      *interface{}
	Value         interface{}
	Values        []interface{}
	ReflectValues []reflect.Value
	EntityValues  []Values
	ValuesV       []Value
)

func SliceToValues(v interface{}) (r Values) {
	valueOfV := reflect.ValueOf(v)
	for i := 0; i < valueOfV.Len(); i++ {
		r = append(r, Value(valueOfV.Index(i).Interface()))
	}
	return r
}
func ToValues(v []Value) (r Values) {
	for _, o := range v {
		r = append(r, o)
	}
	return
}

func (v Values) ValueOfIndex(i int) (r reflect.Value) {
	return reflect.ValueOf(v[i])
}

func (v ReflectValues) ToValueListReflection() (r reflect.Value) {
	a := Values{}
	for _, o := range v {
		a = append(a, o.Interface())
	}
	return reflect.ValueOf(a)
}

func (v ReflectValues) ToValues() (r Values) {
	for _, o := range v {
		r = append(r, o)
	}
	return r

}
