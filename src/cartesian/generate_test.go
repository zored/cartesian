package cartesian

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/fields"
	"github.com/zored/cartesian/src/cartesian/generator"
	"io/ioutil"
	"reflect"
	"testing"
)

type (
	Root struct {
		Int        int
		String     string
		StringList []string
		Other      *Other
	}
	Other struct {
		Id        int
		Bool      bool
		ForeignId int
	}
	Foreign struct {
		Id   int
		UInt uint
	}
)

func TestGenerate(t *testing.T) {
	is := assert.New(t)
	funcResults := generator.FuncResults{{Value: false}, {Value: true}, {Done: true}}
	funcResultsI := -1
	nextForeignId := 0
	ctx := configs.NewContext()
	entities, err := Generate(&configs.Config{
		EntityTemplate: (*Root)(nil),
		Fields: fields.NewFields(
			fields.NewGenerated("Int", generator.NewList(1, 2)),
			fields.NewGenerated("String", generator.NewList("a", "b")),
			fields.NewGenerated("StringList", generator.NewGroup(
				generator.NewMap(
					generator.NewList("c", "d"),
					func(v reflect.Value) reflect.Value {
						return reflect.ValueOf(v.Interface().(string) + "_")
					},
				),
			)),
			fields.NewGenerated("Other", generator.NewEntitySingle(&configs.Config{
				EntityTemplate: (*Other)(nil),
				Fields: fields.NewFields(
					fields.NewGenerated("Bool", generator.NewFunc(func() *generator.FuncResult {
						funcResultsI++
						return funcResults[funcResultsI]
					})),
					fields.NewGenerated("ForeignId", generator.NewMap(generator.NewEntitySingle(&configs.Config{
						EntityTemplate: (*Foreign)(nil),
						Fields: fields.NewFields(
							fields.NewGenerated("UInt", generator.NewList(uint(1))),
						),
						PutEntities: func(_ configs.Context, entities abstract.Entities) {
							for _, v := range entities {
								f := v.(*Foreign)
								is.Equal(0, f.Id)
								nextForeignId++
								f.Id = nextForeignId
							}
						},
					}), func(v reflect.Value) reflect.Value {
						return reflect.ValueOf(v.Interface().(*Foreign).Id)
					})),
				),
				PutEntities: func(ctx configs.Context, entities abstract.Entities) {
					is.Equal(abstract.Entities{
						&Other{ForeignId: 1, Bool: false},
						&Other{ForeignId: 1, Bool: true},
					}, entities)
				},
			})),
		),
	}, ctx)
	is.NoError(err)
	l := []string{"c_", "d_"}
	is.Equal(entities, abstract.Entities{
		&Root{Other: &Other{ForeignId: 1, Bool: false}, StringList: l, String: "a", Int: 1},
		&Root{Other: &Other{ForeignId: 1, Bool: false}, StringList: l, String: "a", Int: 2},
		&Root{Other: &Other{ForeignId: 1, Bool: false}, StringList: l, String: "b", Int: 1},
		&Root{Other: &Other{ForeignId: 1, Bool: false}, StringList: l, String: "b", Int: 2},
		&Root{Other: &Other{ForeignId: 1, Bool: true}, StringList: l, String: "a", Int: 1},
		&Root{Other: &Other{ForeignId: 1, Bool: true}, StringList: l, String: "a", Int: 2},
		&Root{Other: &Other{ForeignId: 1, Bool: true}, StringList: l, String: "b", Int: 1},
		&Root{Other: &Other{ForeignId: 1, Bool: true}, StringList: l, String: "b", Int: 2},
	})

	v := 0
	ctx.EachFactory(func(*configs.TemplateFactory) bool {
		v++
		return false
	})
	is.Equal(3, v)
	is.Equal(
		(*(*(*(*(*ctx.AllResults.Entities)[0].Fields)[3].Entities)[0].Fields)[1].Entities)[0].Value,
		&Foreign{
			Id:   1,
			UInt: 1,
		},
	)

	marshal, err := json.Marshal(ctx)
	is.NoError(err)
	err = ioutil.WriteFile("ctx.json", marshal, 0777)
	is.NoError(err)
}
