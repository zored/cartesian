package tag

import "reflect"

type ITag interface {
	Name() string
	Fill(field reflect.Value)
}
