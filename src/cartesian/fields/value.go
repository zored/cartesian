package fields

import (
	"fmt"
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"reflect"
)

type (
	Value struct {
		field configs.Field
		value abstract.ValuePtr
	}
	Values         []*Value
	ValuesByEntity []Values
)

func NewFieldValue(field configs.Field, v abstract.ValuePtr) *Value {
	return &Value{field: field, value: v}
}

func (v Values) Apply(ctx configs.Context, valueOfEntityPtr reflect.Value) error {
	var err error
	valueOfEntity := valueOfEntityPtr.Elem()
	typeOfEntity := valueOfEntity.Type()
	fieldIndices := map[string]int{}
	for i := 0; i < typeOfEntity.NumField(); i++ {
		fieldIndices[typeOfEntity.Field(i).Name] = i
	}
	for _, fieldValue := range v {
		ctx = ctx.WithField(fieldValue.field)
		fieldName := fieldValue.field.GetName()
		fieldI, ok := fieldIndices[fieldName]
		if !ok {
			return fmt.Errorf(`can't find field %s`, prettyFieldName(typeOfEntity, fieldName))
		}
		field := valueOfEntity.Field(fieldI)

		var valueOfFieldValue reflect.Value
		switch v := (*fieldValue.value).(type) {
		case reflect.Value:
			valueOfFieldValue = v
		case configs.LazyValue:
			valueOfFieldValue, err = v.LazyCreate(ctx)
			if err != nil {
				return err
			}
		default:
			valueOfFieldValue = reflect.ValueOf(v)
		}
		i := valueOfFieldValue.Interface()
		if lazy, ok := i.(configs.LazyValue); ok {
			valueOfFieldValue, err = lazy.LazyCreate(ctx)
			if err != nil {
				return err
			}
		}

		if !field.CanSet() {
			return fmt.Errorf(`can't update field %s`, prettyFieldName(typeOfEntity, fieldName))
		}

		typeOfField := field.Type()

		if valueOfFieldValue.Type().AssignableTo(reflect.TypeOf((abstract.Values)(nil))) {
			switch typeOfField.Kind() {
			case reflect.Slice:
				fallthrough
			case reflect.Array:
				a := reflect.New(typeOfField).Elem()
				typeOfElem := typeOfField.Elem()
				for i := 0; i < valueOfFieldValue.Len(); i++ {
					elemToAppend := valueOfFieldValue.Index(i).Elem()
					if typeOfElem.Kind() == reflect.Ptr && elemToAppend.Kind() != reflect.Ptr {
						elemToAppend = elemToAppend.Convert(typeOfElem.Elem()).Addr()
					} else {
						elemToAppend = elemToAppend.Convert(typeOfElem)
					}
					a = reflect.Append(a, elemToAppend)
				}
				valueOfFieldValue = a
			default:
				return fmt.Errorf(
					"can't assign list of values to non-list %s",
					prettyFieldName(typeOfEntity, fieldName),
				)
			}
		}
		if t := valueOfFieldValue.Type(); !typeOfField.AssignableTo(t) {
			return fmt.Errorf(
				`"%s" is not assignable to to %s (type "%s")`,
				t.Name(),
				prettyFieldName(typeOfEntity, fieldName),
				typeOfField,
			)
		}
		ctx = ctx.WithFieldValuePointer(valueOfFieldValue.Interface())
		field.Set(valueOfFieldValue)
	}
	return nil
}

func prettyFieldName(typeOfEntity reflect.Type, fieldName string) string {
	return fmt.Sprintf(`"%s.%s"`, typeOfEntity.Name(), fieldName)
}
