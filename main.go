package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	deck := Deck{}
	deck.setDeck()
	fmt.Println(deck.cards[0])
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("プレイヤーの名前, 所持金を記入してください")
	sc.Scan()
	t := strings.Split(sc.Text(), ",")
	name := ""
	money := 0
	for i, s := range t {
		if i == 0 {
			name += s
		} else {
			m, _ := strconv.Atoi(s)
			money += m
		}
	}
	player := Player{name, money}
	fmt.Println(player)

	fmt.Println("ディーラーの名前, 所持金を記入してください。")
	sc.Scan()
	t = strings.Split(sc.Text(), ", ")
	name = ""
	money = 0
	for i, s := range t {
		if i == 0 {
			name += s
		} else {
			m, _ := strconv.Atoi(s)
			money += m
		}
	}
	dealer := Dealer{name, money}
	fmt.Println(dealer)
}
