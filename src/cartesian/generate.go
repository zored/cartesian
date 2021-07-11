package cartesian

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/fields"
	"reflect"
)

func Generate(c *configs.Config) (r abstract.Entities, err error) {
	ios := c.Flatten(true)
	ctx := &configs.Context{}
	for _, io := range ios {
		config := io.GetInput()
		entities := io.GetOutput()
		for _, values := range getValuesByEntity(ctx, config.Fields) {
			entity, err := createEntity(config.EntityTemplate, values)
			if err != nil {
				return nil, err
			}
			entities = append(entities, entity)
		}
		io.SetOutput(entities)

		//prevIo := ctx.GetLastCompleteIO()

		ctx.AddCompleteIO(io)
		if config.PutIO != nil {
			config.PutIO(io)
		}

		//if prevIo == nil {
		//	continue
		//}
		//
		//fs := prevIo.GetInput().Fields
		//for i := 0; i < fs.Len(); i++ {
		//	var parent abstract.Value
		//	field := fs.Index(i)
		//	fields.NewFieldValue(field, abstract.ValueAddr(field.GetParentValue(parent)))
		//}
	}
	if l := ios.Last(); l != nil {
		r = l.GetOutput()
	}
	return
}

func createEntity(tmpl configs.EntityTemplate, values fields.Values) (abstract.Entity, error) {
	entity := reflect.New(reflect.TypeOf(tmpl).Elem())
	if err := values.Apply(entity); err != nil {
		return nil, err
	}
	return entity.Interface(), nil
}

func getValuesByEntity(ctx *configs.Context, fieldList configs.Fields) (r fields.ValuesByEntity) {
	type intByFieldIndex map[int]int

	valuesByFieldIndex := map[int]abstract.Values{}
	lens := intByFieldIndex{}
	valueIndices := intByFieldIndex{}
	for i, v := range fieldList.CreateEntityValues(ctx) {
		valuesByFieldIndex[i] = v
		lens[i] = len(v)
		valueIndices[i] = 0
	}

	for {
		// Create entity values:
		lenFields := fieldList.Len()
		lastFieldI := lenFields - 1
		v := fields.Values{}
		for fieldI := 0; fieldI < lenFields; fieldI++ {
			f := fieldList.Index(fieldI)
			values := valuesByFieldIndex[fieldI]
			v = append(v, fields.NewFieldValue(f, &values[valueIndices[fieldI]]))
		}
		r = append(r, v)

		// Increment index:
		for fieldI := 0; fieldI < lenFields; fieldI++ {
			valueIndices[fieldI]++
			if valueIndices[fieldI] < lens[fieldI] {
				break
			}
			if fieldI == lastFieldI {
				return r
			}
			valueIndices[fieldI] = 0
		}
	}
}
