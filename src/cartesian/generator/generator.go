package generator

import "reflect"

type Generator interface {
	Next() reflect.Value
	Done() bool
}
