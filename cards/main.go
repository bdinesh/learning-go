package main

func main() {
	cards := newDeck()

	// deckAtHand, remainingDeck := cards.deal(5)

	// deckAtHand.print()
	// remainingDeck.print()

	// cards.saveToFile("mydeck")
	// cards = newDeckFromFile("mydeck")
	cards.shuffle()
	cards.print()
}
