package configs

import "github.com/zored/cartesian/src/cartesian/abstract"

type (
	TemplateFactory struct {
		IO     IO
		create TemplateFactoryF
	}
	TemplateFactoryF  func(Context, *Config) (abstract.Entities, error)
	TemplateFactories []*TemplateFactory
)

func (f TemplateFactories) Prepend(v *TemplateFactory) TemplateFactories {
	return append(TemplateFactories{v}, f...)
}

func (f TemplateFactories) First() *TemplateFactory {
	if len(f) == 0 {
		return nil
	}
	return f[0]
}

func NewTemplateFactory(io IO, create TemplateFactoryF) *TemplateFactory {
	return &TemplateFactory{IO: io, create: create}
}

func (f *TemplateFactory) Create(ctx Context) (abstract.Entities, error) {
	if f == nil {
		return nil, nil
	}
	io := f.IO
	r, err := f.create(ctx, io.GetInput())
	if err != nil {
		return nil, err
	}
	return r, nil
}
