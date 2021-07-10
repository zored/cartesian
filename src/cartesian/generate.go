package cartesian

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/generator"
	"github.com/zored/cartesian/src/cartesian/tag"
	"reflect"
)

// Entity is pointer on final entity.
type Entity interface{}

func Generate(c *Config) (r abstract.Entities, err error) {
	fillTagsValues(c.Tags)
	for _, values := range getValuesByEntity(c.Tags) {
		e, err := createEntity(c.EntityTemplate, values)
		if err != nil {
			return nil, err
		}
		r = append(r, e)
	}
	return r, err
}

func createEntity(tmpl EntityTemplate, values tag.Values) (Entity, error) {
	entity := reflect.New(reflect.TypeOf(tmpl).Elem())
	if err := values.Apply(entity); err != nil {
		return nil, err
	}
	return entity.Interface(), nil
}

func getValuesByEntity(tags tag.Tags) (r tag.ValuesByEntity) {
	type intByTagIndex map[int]int
	lens := intByTagIndex{}
	valueIndices := intByTagIndex{}

	for i, t := range tags {
		lens[i] = len(t.Values)
		valueIndices[i] = 0
	}

	for {
		// Create entity values:
		lenTags := len(tags)
		lastTagI := lenTags - 1
		v := tag.Values{}
		for tagI := 0; tagI < lenTags; tagI++ {
			t := tags[tagI]
			v = append(v, tag.NewTagValue(t, &t.Values[valueIndices[tagI]]))
		}
		r = append(r, v)

		// Increment index:
		for tagI := 0; tagI < lenTags; tagI++ {
			valueIndices[tagI]++
			if valueIndices[tagI] < lens[tagI] {
				break
			}
			if tagI == lastTagI {
				return r
			}
			valueIndices[tagI] = 0
		}
	}
}

func fillTagsValues(tags tag.Tags) {
	for _, t := range tags {
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
