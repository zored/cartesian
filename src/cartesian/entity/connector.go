package entity

import "github.com/zored/cartesian/src/cartesian/abstract"

type Count int

const (
	One Count = iota + 1
	Many
)

type Connector interface {
	OneToOne(a Definition, b Definition, connect func(a, b abstract.Instance))
	OneToMany(one Definition, many Definition)
	Instantiate() map[Definition]abstract.Instances
}

// Worker <- Company
// Unit <- Company
// Worker <- Company
