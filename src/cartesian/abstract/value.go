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

func ValueAddr(v Value) ValuePtr {
	if v == nil {
		return nil
	}
	vi := v.(interface{})
	return &vi
}

func ToValues(v []Value) (r Values) {
	for _, o := range v {
		r = append(r, o)
	}
	return
}

func ListOfOne(v Value) Values {
	return Values{v}
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
