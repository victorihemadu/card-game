package main

import (
	//"fmt"
	"github.com/bigDev/cardsGame/deck"
)




func main() {
	cards := deck.NewDeckFromFile("my_cards")

	cards.Print()
	
	// cards.Print()
	
	// hand, remainingCards := deck.Deal(cards, 5)

	// hand.Print()
	// remainingCards.Print()


}
