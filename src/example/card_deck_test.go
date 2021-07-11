package example

import (
	"fmt"
	. "github.com/zored/cartesian/src/cartesian"
	"github.com/zored/cartesian/src/cartesian/abstract"
	"github.com/zored/cartesian/src/cartesian/configs"
	"github.com/zored/cartesian/src/cartesian/fields"
	"github.com/zored/cartesian/src/cartesian/generator"
	"github.com/zored/cartesian/src/example/entities"
	"reflect"
)

var (
	cardFields = fields.NewFields(
		fields.NewGenerated("Rank", generator.NewList("Ace", "King", "Queen", "Jack", "10", "9", "8", "7", "6", "5", "4", "3", "2")),
		fields.NewGenerated("Suit", generator.NewList("diamonds", "clubs", "hearts", "spades")),
	)
	cardConfig = &configs.Config{
		EntityTemplate: (*entities.Card)(nil),
		Fields:         cardFields,
	}
)

// ExampleGenerateDeck generates 52 cards based on 13 ranks and 4 suits.
func ExampleGenerateDeck() {
	deck := entities.Deck{}
	r, err := Generate(cardConfig)
	panicOnErr(err)
	r.Each(func(v abstract.Entity) {
		deck = append(deck, v.(*entities.Card))
	})
	fmt.Printf("%v", deck)
	// Output: Ace of diamonds, King of diamonds, Queen of diamonds, Jack of diamonds, 10 of diamonds, 9 of diamonds, 8 of diamonds, 7 of diamonds, 6 of diamonds, 5 of diamonds, 4 of diamonds, 3 of diamonds, 2 of diamonds, Ace of clubs, King of clubs, Queen of clubs, Jack of clubs, 10 of clubs, 9 of clubs, 8 of clubs, 7 of clubs, 6 of clubs, 5 of clubs, 4 of clubs, 3 of clubs, 2 of clubs, Ace of hearts, King of hearts, Queen of hearts, Jack of hearts, 10 of hearts, 9 of hearts, 8 of hearts, 7 of hearts, 6 of hearts, 5 of hearts, 4 of hearts, 3 of hearts, 2 of hearts, Ace of spades, King of spades, Queen of spades, Jack of spades, 10 of spades, 9 of spades, 8 of spades, 7 of spades, 6 of spades, 5 of spades, 4 of spades, 3 of spades, 2 of spades
}

var (
	croupierFields = fields.NewFields(
		fields.NewGenerated("Name", generator.NewList("Bob", "Rob")),
		fields.NewGenerated("Deck", generator.NewEntityList(cardConfig)),
	)
	croupierConfig = &configs.Config{
		EntityTemplate: (*entities.Croupier)(nil),
		Fields:         croupierFields,
	}
)

func ExampleGenerateCroupier() {
	croupiers := entities.Croupiers{}
	r, err := Generate(croupierConfig)
	panicOnErr(err)
	r.Each(func(v abstract.Entity) {
		croupiers = append(croupiers, v.(*entities.Croupier))
	})
	fmt.Printf("%v", croupiers)
	// Output: [Bob (Ace of diamonds, King of diamonds, Queen of diamonds, Jack of diamonds, 10 of diamonds, 9 of diamonds, 8 of diamonds, 7 of diamonds, 6 of diamonds, 5 of diamonds, 4 of diamonds, 3 of diamonds, 2 of diamonds, Ace of clubs, King of clubs, Queen of clubs, Jack of clubs, 10 of clubs, 9 of clubs, 8 of clubs, 7 of clubs, 6 of clubs, 5 of clubs, 4 of clubs, 3 of clubs, 2 of clubs, Ace of hearts, King of hearts, Queen of hearts, Jack of hearts, 10 of hearts, 9 of hearts, 8 of hearts, 7 of hearts, 6 of hearts, 5 of hearts, 4 of hearts, 3 of hearts, 2 of hearts, Ace of spades, King of spades, Queen of spades, Jack of spades, 10 of spades, 9 of spades, 8 of spades, 7 of spades, 6 of spades, 5 of spades, 4 of spades, 3 of spades, 2 of spades) Rob (Ace of diamonds, King of diamonds, Queen of diamonds, Jack of diamonds, 10 of diamonds, 9 of diamonds, 8 of diamonds, 7 of diamonds, 6 of diamonds, 5 of diamonds, 4 of diamonds, 3 of diamonds, 2 of diamonds, Ace of clubs, King of clubs, Queen of clubs, Jack of clubs, 10 of clubs, 9 of clubs, 8 of clubs, 7 of clubs, 6 of clubs, 5 of clubs, 4 of clubs, 3 of clubs, 2 of clubs, Ace of hearts, King of hearts, Queen of hearts, Jack of hearts, 10 of hearts, 9 of hearts, 8 of hearts, 7 of hearts, 6 of hearts, 5 of hearts, 4 of hearts, 3 of hearts, 2 of hearts, Ace of spades, King of spades, Queen of spades, Jack of spades, 10 of spades, 9 of spades, 8 of spades, 7 of spades, 6 of spades, 5 of spades, 4 of spades, 3 of spades, 2 of spades)]
}

func ExampleGenerateCasino() {
	var (
		croupierFields = append(
			croupierFields,
			//fields.NewFromParent("CasinoId", 0, func(p abstract.Value) (fieldValue abstract.Value) {
			//	return p.(*entities.Casino).Id
			//}),
		)
		lastCroupierId = 0
		croupiers      = entities.Croupiers{}
		croupierConfig = &configs.Config{
			EntityTemplate: (*entities.Croupier)(nil),
			Fields:         croupierFields,
			PutIO: func(io configs.IO) {
				io.GetOutput().Each(func(e abstract.Entity) {
					c := e.(*entities.Croupier)
					if c.Id > 0 {
						return
					}
					lastCroupierId++
					c.Id = lastCroupierId
					croupiers = append(croupiers, c)
				})
			},
		}
		casinoFields = fields.NewFields(
			fields.NewGenerated(
				"Name",
				generator.NewList("Las Vegas", "Super Slots"),
			),
			fields.NewGenerated(
				"CroupierIds",
				generator.NewGroup(generator.NewMap(
					generator.NewEntitySingle(croupierConfig),
					func(v reflect.Value) reflect.Value {
						return reflect.ValueOf(v.Interface().(*entities.Croupier).Id)
					},
				)),
			),
		)
		lastCasinoId = 0
		casinoConfig = &configs.Config{
			EntityTemplate: (*entities.Casino)(nil),
			Fields:         casinoFields,
			PutIO: func(io configs.IO) {
				io.GetOutput().Each(func(e abstract.Entity) {
					c := e.(*entities.Casino)
					if c.Id > 0 {
						return
					}
					lastCasinoId++
					c.Id = lastCasinoId
				})
			},
		}
	)
	casinos := []*entities.Casino{}
	r, err := Generate(casinoConfig)
	panicOnErr(err)
	r.Each(func(v abstract.Entity) {
		casinos = append(casinos, v.(*entities.Casino))
	})
	fmt.Printf("%v, %s", casinos, croupiers.ShortString())
	// Output: [casino #1 with croupier IDs: [1 2] casino #2 with croupier IDs: [1 2]], #1 Bob with 52 cards, #2 Rob with 52 cards
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
