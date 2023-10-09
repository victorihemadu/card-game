package deck

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Deck []string

func NewDeck() Deck {
	cards := Deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+ " of " +suit)
		}
	}
	return cards
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Deal(d Deck, handsize int) (Deck, Deck) {
	return d[:handsize], d[handsize:]
}

func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) SaveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	 if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return Deck(s)
	
}