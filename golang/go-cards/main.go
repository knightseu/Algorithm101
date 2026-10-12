package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Diamond"
	// card = "Five of Heart"
	card := newCard()
	fmt.Println(card)
}
func newCard() string {
	return "Three of Club"
}

// Basic Go type
// bool, string, int, float64
