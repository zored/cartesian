package cartesian

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/fields"
	"github.com/zored/cartesian/src/cartesian/generator"
	"reflect"
)

// Entity is pointer on final entity.
type Entity interface{}

func Generate(c *Config) (r abstract.Entities, err error) {
	fillFieldsValues(c.Fields)
	for _, values := range getValuesByEntity(c.Fields) {
		e, err := createEntity(c.EntityTemplate, values)
		if err != nil {
			return nil, err
		}
		r = append(r, e)
	}
	return r, err
}

func createEntity(tmpl EntityTemplate, values fields.Values) (Entity, error) {
	entity := reflect.New(reflect.TypeOf(tmpl).Elem())
	if err := values.Apply(entity); err != nil {
		return nil, err
	}
	return entity.Interface(), nil
}

func getValuesByEntity(fs fields.Fields) (r fields.ValuesByEntity) {
	type intByFieldIndex map[int]int
	lens := intByFieldIndex{}
	valueIndices := intByFieldIndex{}
	for i, t := range fs {
		lens[i] = len(t.Values)
		valueIndices[i] = 0
	}

	for {
		// Create entity values:
		lenFields := len(fs)
		lastFieldI := lenFields - 1
		v := fields.Values{}
		for fieldI := 0; fieldI < lenFields; fieldI++ {
			t := fs[fieldI]
			v = append(v, fields.NewFieldValue(t, &t.Values[valueIndices[fieldI]]))
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

func fillFieldsValues(fields fields.Fields) {
	for _, t := range fields {
		if g := t.Generator; g != nil {
			t.Values = generateValues(g)
		}
	}
}

func generateValues(g generator.Generator) (r abstract.Values) {
	for !g.Done() {
		r = append(r, g.Next())
	}
	return r
}
