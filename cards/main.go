package main

// func main() {
// 	//var card string = "Ace of Spades"
// 	card := "Ace of Spades"
// 	card = "Five of Diamonds"

// 	fmt.Println(card)

// }
// func main() {
// 	card := newCard()

// 	fmt.Println(card)
// }

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
