package configs

import "github.com/zored/cartesian/src/cartesian/abstract"

type (
	ResultVisitor struct {
		Entities *ResultEntities
	}
	LocalResult struct {
		entity   *ResultEntity
		field    *ResultField
		entities *ResultEntities
	}
)

func NewResultVisitor() (*ResultVisitor, LocalResult) {
	r := &ResultEntities{}
	return &ResultVisitor{Entities: r}, LocalResult{entities: r}
}

func (l *LocalResult) SetConfig(config *Config) {
	config.FillName()
	l.entity = &ResultEntity{
		Config: config,
		Fields: &ResultFields{},
		Value:  nil,
	}
	*l.entities = append(*l.entities, l.entity)
}

func (l *LocalResult) SetEntity(v abstract.Instance) {
	if l.entity.valueSet {
		l.SetConfig(l.entity.Config)
	}
	l.entity.valueSet = true
	l.entity.Value = v
}

func (l *LocalResult) SetField(field Field) {
	l.entities = &ResultEntities{}
	l.field = &ResultField{
		Value:    nil,
		Config:   field,
		Entities: l.entities,
	}
	*l.entity.Fields = append(*l.entity.Fields, l.field)
}

func (l *LocalResult) SetFieldValuePointer(valuePtr interface{}) {
	if l.field.valueSet {
		l.SetField(l.field.Config)
	}
	l.field.valueSet = true
	l.field.Value = valuePtr
}
