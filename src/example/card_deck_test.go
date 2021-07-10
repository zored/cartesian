package example

import (
	"fmt"
	. "github.com/zored/cartesian/src/cartesian"
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/fields"
	"github.com/zored/cartesian/src/cartesian/generator"
)

type Card struct {
	Rank string
	Suit string
}

type Cards []*Card

func (c *Card) String() string {
	return fmt.Sprintf("%s %s", c.Rank, c.Suit)
}
func ExampleGenerate() {
	cards := Cards{}
	r, err := Generate(&Config{
		EntityTemplate: (*Card)(nil),
		Fields: fields.NewFields(
			// You can generate:
			fields.NewGenerated(
				"Rank",
				generator.NewSet("A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"),
			),
			fields.NewValued("Suit", "diamonds", "clubs", "hearts", "spades"),
		),
	})
	if err != nil {
		panic(err)
	}
	r.EachEntity(func(v abstract.Entity) {
		cards = append(cards, v.(*Card))
	})

	fmt.Printf("%v", cards)
	// Output: [A diamonds K diamonds Q diamonds J diamonds 10 diamonds 9 diamonds 8 diamonds 7 diamonds 6 diamonds 5 diamonds 4 diamonds 3 diamonds 2 diamonds A clubs K clubs Q clubs J clubs 10 clubs 9 clubs 8 clubs 7 clubs 6 clubs 5 clubs 4 clubs 3 clubs 2 clubs A hearts K hearts Q hearts J hearts 10 hearts 9 hearts 8 hearts 7 hearts 6 hearts 5 hearts 4 hearts 3 hearts 2 hearts A spades K spades Q spades J spades 10 spades 9 spades 8 spades 7 spades 6 spades 5 spades 4 spades 3 spades 2 spades]
}
