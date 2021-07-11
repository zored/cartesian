package entities

import (
	"fmt"
	"strings"
)

type (
	Card struct {
		Rank string
		Suit string
	}
	Deck     []*Card
	Croupier struct {
		CasinoId int
		Name     string
		Deck     Deck
		Id       int
	}
	Croupiers []*Croupier
	Casino    struct {
		Name        string
		CroupierIds []int
		Id          int
	}
)

func (c *Card) String() string {
	return fmt.Sprintf("%s %s", c.Rank, c.Suit)
}

func (c Deck) String() (s string) {
	for _, v := range c {
		s += v.Rank + " of " + v.Suit + ", "
	}
	return strings.TrimRight(s, ", ")
}

func (c *Croupier) String() string {
	return fmt.Sprintf("%s (%s)", c.Name, c.Deck)
}

func (c Croupiers) ShortString() string {
	r := ""
	for _, v := range c {
		r += fmt.Sprintf("#%d %s with %d cards, ", v.Id, v.Name, len(v.Deck))
	}
	return strings.TrimRight(r, ", ")
}

func (c *Casino) String() (s string) {
	return fmt.Sprintf("casino #%d with croupier IDs: %v", c.Id, c.CroupierIds)
}
