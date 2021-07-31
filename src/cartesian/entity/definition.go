package entity

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
)

type (
	Definition interface {
		Variants() int
		Instantiate() abstract.Instances
	}
	DefinitionApplication struct {
		definition Definition
		Reason string
	}
	definition struct {
		fields configs.Fields
	}
)

