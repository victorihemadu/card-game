package main

import (
	"github.com/bigDev/cardsGame/deck"
)


func main() {
	cards := deck.NewDeck()
	
	cards.Print()
}
