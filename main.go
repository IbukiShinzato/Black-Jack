package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	deck := Deck{}
	deck.setDeck()

	player, dealer := start()
	game(deck, player, dealer)
}

func start() (Player, Dealer) {
	fmt.Printf("\nーーーブラックジャック！ーーー\n\n")
	sc := bufio.NewScanner(os.Stdin)
	var hand []Card

	fmt.Printf("\nプレイヤーの名前を入力してください。\n")
	sc.Scan()
	name := sc.Text()

	fmt.Printf("\nプレイヤーの所持金を入力してください。\n")
	sc.Scan()
	money, _ := strconv.Atoi(sc.Text())
	player := Player{name, money, hand}

	fmt.Printf("\nディーラーの名前を入力してください。\n")
	sc.Scan()
	name = sc.Text()

	dealer := Dealer{name, hand}

	return player, dealer
}

func game(deck Deck, player Player, dealer Dealer) {
	fmt.Printf("\nーーーゲーム開始！ーーー\n\n")
	deck.shuffle()
	player.hand = []Card{}
	dealer.hand = []Card{}

	cnt := 0
	player.hand = append(player.hand, deck.cards[cnt])
	cnt += 1
	player.hand = append(player.hand, deck.cards[cnt])
	cnt += 1
	dealer.hand = append(dealer.hand, deck.cards[cnt])
	cnt += 1
	dealer.hand = append(dealer.hand, deck.cards[cnt])
}
