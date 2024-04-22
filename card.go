package main

import (
	"math/rand"
	"strconv"
	"time"
)

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

func (d *Deck) shuffle() {
	var n int
	rand.Seed(time.Now().UnixNano())
	cards := make([]Card, 52)
	m := make(map[int]bool)
	i := 0
	for i < 52 {
		n = rand.Intn(52)
		if !m[n] {
			cards[i] = d.cards[n]
			i += 1
			m[n] = true
		}
	}
	d.cards = cards
}
