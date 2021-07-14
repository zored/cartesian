package cartesian

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/fields"
	"reflect"
)

func Generate(c *configs.Config, ctxs ...configs.Context) (abstract.Entities, error) {
	ios := c.Flatten(true)
	ctx := configs.NewContext()
	for _, c := range ctxs {
		ctx = c
	}
	factories := configs.TemplateFactories{}
	for _, io := range ios {
		io := io
		factories = factories.Prepend(configs.NewTemplateFactory(
			io,
			// todo move out logic
			func(ctx configs.Context, config *configs.Config) (r abstract.Entities, err error) {
				ctx = ctx.WithConfig(config)
				byEntity, err := getValuesByEntity(ctx, config.Fields)
				if err != nil {
					return nil, err
				}
				for _, values := range byEntity {
					entity, err := createEntity(ctx, config.EntityTemplate, values)
					if err != nil {
						return nil, err
					}
					r = append(r, entity)
				}
				if config.PutEntities != nil {
					config.PutEntities(ctx, r)
				}
				return r, nil
			},
		))
		factories = append(configs.TemplateFactories{}, factories...)
	}
	ctx = ctx.WithFactories(factories)
	return factories.First().Create(ctx)
}

func createEntity(ctx configs.Context, tmpl configs.EntityTemplate, values fields.Values) (abstract.Entity, error) {
	entity := reflect.New(reflect.TypeOf(tmpl).Elem())
	r := abstract.Entity(entity.Interface())
	ctx = ctx.WithEntity(r)
	if err := values.Apply(ctx, entity); err != nil {
		return nil, err
	}
	return r, nil
}

func getValuesByEntity(ctx configs.Context, fieldList configs.Fields) (r fields.ValuesByEntity, err error) {
	type intByFieldIndex map[int]int

	valuesByFieldIndex := map[int]abstract.Values{}
	lens := intByFieldIndex{}
	valueIndices := intByFieldIndex{}
	values, err := fieldList.CreateEntityValues(ctx)
	if err != nil {
		return nil, err
	}
	for i, v := range values {
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
				return r, err
			}
			valueIndices[fieldI] = 0
		}
	}
}
