package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Diamond"
	// card = "Five of Heart"
	// card := newCard()
	// fmt.Println(card)

	/////
	// cards := []string{newCard(), "Ace of Diamond"} //slice
	// fmt.Println((cards))                           // [Three of Club Ace of Diamond]
	// cards = append(cards, "Six of Spades")
	// fmt.Println((cards)) //[Three of Club Ace of Diamond Six of Spades]

	// fmt.Println("---")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// fmt.Println("===")
	// cards := deck{newCard(), "Ace of Diamond"}
	// cards.print()

	// fmt.Println("---")
	cards := newDeck()
	cards.print()
	// hands, remainingCards := deal(cards, 5)
	// hands.print()
	// fmt.Println("***")
	// remainingCards.print()

	//
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")

	cards = newDeckFromFile("my_cards")
	cards.print()
}

func newCard() string {
	return "Three of Club"
}

// Basic Go type
// bool, string, int, float64
