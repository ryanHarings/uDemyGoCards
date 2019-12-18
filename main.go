package main

func main() {
	// var card string
	//	var card string = "Ace of Spades"
	// card := newCard()
	// cards := deck{newCard(), newCard(), "Queen of Hearts"}
	// cards = append(cards, "Six of Spades")
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
