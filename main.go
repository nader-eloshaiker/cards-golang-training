package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remaingDeck := deal(cards, 5)

	hand.print()
	remaingDeck.print()

	fmt.Println(hand.toString())
}