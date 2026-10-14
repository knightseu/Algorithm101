package main

import (
	"fmt"
	"io/ioutil"
	// "ioutil"
	"os"
	"strings"
)

// create a new type of 'deck'
// which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Hearts", "Club"}
	cardValues := []string{"Ace", "2", "3", "4"}
	for _, suit := range cardSuits {
		for _, v := range cardValues {
			cards = append(cards, v+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") //pkg: strings
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//
		fmt.Println("Error read from file: ", err)
		os.Exit(1)
		// return newDeck()
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
