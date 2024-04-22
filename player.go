package main

type Player struct {
	name  string
	money int
	hand  []Card
}

type Dealer struct {
	name string
	hand []Card
}

// func (d *Dealer) shuffle(p Player) {
// 	p.deck = append(p.deck)
// 	d.deck = append(d.deck)
// }
