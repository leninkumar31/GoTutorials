package main

func main() {
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// cards := newDeckFromFile("deckOfCardsFile")
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards.saveToFile("deckOfCardsFile")
}
