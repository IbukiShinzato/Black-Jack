package main

import "strconv"

type Card struct {
	suit string
	num  int
}

type Deck struct {
	cards []Card
}

func (c *Card) getCard() string {
	return c.suit + strconv.Itoa(c.num)
}

func (d *Deck) setDeck() {
	suits := []string{"H", "D", "K", "S"}
	for i := 1; i < 14; i++ {
		for _, s := range suits {
			c := Card{s, i}
			d.cards = append(d.cards, c)
		}
	}
}
