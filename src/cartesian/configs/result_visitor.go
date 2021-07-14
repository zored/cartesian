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

func (c *ResultVisitor) SetConfig(config *Config, l *LocalResult) {
	config.FillName()
	l.entity = &ResultEntity{
		Config: config,
		Fields: &ResultFields{},
		Value:  nil,
	}
	*l.entities = append(*l.entities, l.entity)
}

func (c *ResultVisitor) SetEntity(v abstract.Entity, l *LocalResult) {
	if l.entity.valueSet {
		c.SetConfig(l.entity.Config, l)
	}
	l.entity.valueSet = true
	l.entity.Value = v
}

func (c *ResultVisitor) SetField(field Field, l *LocalResult) {
	l.entities = &ResultEntities{}
	l.field = &ResultField{
		Value:    nil,
		Config:   field,
		Entities: l.entities,
	}
	*l.entity.Fields = append(*l.entity.Fields, l.field)
}

func (c *ResultVisitor) SetFieldValuePointer(valuePtr interface{}, l *LocalResult) {
	if l.field.valueSet {
		c.SetField(l.field.Config, l)
	}
	l.field.valueSet = true
	l.field.Value = valuePtr
}
