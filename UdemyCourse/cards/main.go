package main

func main() {
	// when a variable is initialized I can use :=
	// var card string = "Ace of Spades"
	// slice is a dynamic array
	cards := newDeck()
	cards.shuffle()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	cards.print()
}
