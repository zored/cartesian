package configs

import "github.com/zored/cartesian/src/cartesian/abstract"

type (
	TemplateFactory struct {
		IO     IO
		create TemplateFactoryF
	}
	ITemplateFactory interface {
		Create(ctx Context) (abstract.Instances, error)
	}
	TemplateFactoryF  func(Context, *Config) (abstract.Instances, error)
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

func (f *TemplateFactory) Create(ctx Context) (abstract.Instances, error) {
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
