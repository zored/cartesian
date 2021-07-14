package configs

import (
	"reflect"
)

type LazyValue interface {
	LazyCreate(ctx Context) (reflect.Value, error)
}
