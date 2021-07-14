package configs

import (
	"github.com/zored/cartesian/src/cartesian/abstract"
)

type (
	ResultEntity struct {
		Value    abstract.Entity
		Config   *Config
		Fields   *ResultFields
		valueSet bool
	}
	ResultField struct {
		Config   Field `json:",inline"`
		Value    abstract.Value
		Entities *ResultEntities `json:",omitempty"`
		valueSet bool
	}
	ResultEntities []*ResultEntity
	ResultFields   []*ResultField
)
