package cartesian

import (
	"github.com/stretchr/testify/assert"
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/fields"
	"github.com/zored/cartesian/src/cartesian/generator"
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
						PutIO: func(io configs.IO) {
							io.GetOutput().Each(func(v abstract.Entity) {
								f := v.(*Foreign)
								is.Equal(0, f.Id)
								nextForeignId++
								f.Id = nextForeignId
							})
						},
					}), func(v reflect.Value) reflect.Value {
						return reflect.ValueOf(v.Interface().(*Foreign).Id)
					})),
				),
				PutIO: func(io configs.IO) {
					is.Equal(abstract.Entities{
						&Other{ForeignId: 1, Bool: false},
						&Other{ForeignId: 1, Bool: true},
					}, io.GetOutput())
				},
			})),
		),
	})
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
}
