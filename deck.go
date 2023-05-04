package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suits := range cardSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+"Of"+suits)
		}
	}
	return cards
}

func (d deck) print() {
	for _, cards := range d {
		log(cards)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

// func (d deck) readFromFile(filename string) ([]byte, error) {
// 	return ioutil.ReadFile(filename)
// }

// func (d deck) turnByteToDeck(s []byte) []string {
// 	deck := string(s)
// 	return strings.Split(deck, ",")

// }
