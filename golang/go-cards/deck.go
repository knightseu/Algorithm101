package main

import "fmt"

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

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
