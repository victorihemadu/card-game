package main

import (
	"fmt"
	"github.com/bigDev/cardsGame/deck"
)




func main() {
	cards := deck.NewDeck()

	fmt.Println(cards.ToString())
	
	// cards.Print()
	
	// hand, remainingCards := deck.Deal(cards, 5)

	// hand.Print()
	// remainingCards.Print()


}
