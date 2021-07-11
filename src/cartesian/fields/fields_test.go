package fields

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	mock_config "github.com/zored/cartesian/src/cartesian/configs/mocks"
	mock_generator "github.com/zored/cartesian/src/cartesian/generator/mocks"
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	is := assert.New(t)
	generator1 := mock_generator.NewMockGenerator(ctrl)
	generator2 := mock_generator.NewMockGenerator(ctrl)
	field1 := NewGenerated("field1", generator1)
	field2 := NewGenerated("field2", generator2)
	fields := NewFields(field1, field2)
	is.Equal(2, fields.Len())
	is.Nil(fields.Index(-2))
	is.Nil(fields.Index(+2))
	is.Equal(field1, fields.Index(0))
	is.Equal(field2, fields.Index(1))
	io1 := mock_config.NewMockIO(ctrl)
	io2 := mock_config.NewMockIO(ctrl)
	ios1 := configs.IOs{io1}
	ios2 := configs.IOs{io2}
	generator1.EXPECT().GetIOs().Return(ios1)
	generator2.EXPECT().GetIOs().Return(ios2)
	is.Equal(configs.IOs{io1, io2}, fields.GetIOs())
	ctx := configs.NewContext()
	mockGenerator(generator1, ctx, 1)
	mockGenerator(generator2, ctx, 2)
	is.Equal(
		abstract.EntityValues{{reflect.ValueOf(1)}, {reflect.ValueOf(2)}},
		fields.CreateEntityValues(ctx),
	)
}

func mockGenerator(generator *mock_generator.MockGenerator, ctx *configs.Context, value interface{}) {
	state := "any"
	generator.EXPECT().State(ctx).Return(state)
	generator.EXPECT().Done(state).Return(false)
	generator.EXPECT().Next(state).Return(reflect.ValueOf(value))
	generator.EXPECT().Done(state).Return(true)
}
