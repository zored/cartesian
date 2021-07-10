package fields

import (
	"fmt"
	"github.com/zored/cartesian/src/cartesian/abstract"
	"reflect"
)

type (
	Value struct {
		field   *field
		value abstract.ValuePtr
	}
	Values         []*Value
	ValuesByEntity []Values
)

func NewFieldValue(t *field, v abstract.ValuePtr) *Value {
	return &Value{field: t, value: v}
}

func (v Values) Apply(valueOfEntityPtr reflect.Value) error {
	valueOfEntity := valueOfEntityPtr.Elem()
	typeOfEntity := valueOfEntity.Type()
	fieldIndices := map[string]int{}
	for i := 0; i < typeOfEntity.NumField(); i++ {
		fieldIndices[typeOfEntity.Field(i).Name] = i
	}
	for _, fieldValue := range v {
		fieldName := fieldValue.field.Name
		fieldI, ok := fieldIndices[fieldName]
		if !ok {
			return fmt.Errorf(`can't find field %s`, prettyFieldName(typeOfEntity, fieldName))
		}
		field := valueOfEntity.Field(fieldI)

		var valueOfFieldValue reflect.Value
		switch v := (*fieldValue.value).(type) {
		case reflect.Value:
			valueOfFieldValue = v
		default:
			valueOfFieldValue = reflect.ValueOf(v)
		}

		if !field.CanSet() {
			return fmt.Errorf(`can't update field %s`, prettyFieldName(typeOfEntity, fieldName))
		}
		if t := valueOfFieldValue.Type(); !field.Type().AssignableTo(t) {
			return fmt.Errorf(
				`"%s" is not assignable to to %s (type "%s")`,
				t.Name(),
				prettyFieldName(typeOfEntity, fieldName),
				field.Type(),
			)
		}
		field.Set(valueOfFieldValue)
	}
	return nil
}

func prettyFieldName(typeOfEntity reflect.Type, fieldName string) string {
	return fmt.Sprintf(`"%s.%s"`, typeOfEntity.Name(), fieldName)
}
