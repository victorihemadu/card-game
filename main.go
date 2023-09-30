package main

import (
	"fmt"
	"github.com/bigDev/cardsGame/deck"
)


func main() {
	cards := deck.Deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}