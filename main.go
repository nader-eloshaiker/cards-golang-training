package main

func main() {
	cards := newDeck()

	hand, remaingDeck := deal(cards, 5)

	hand.print()
	remaingDeck.print()
}