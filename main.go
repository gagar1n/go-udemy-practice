package main

func main() {
	cards := newDeck()
	cards.saveToFile("my.cards")

	cards2 := newDeckFromFile("my.cards")
	cards2.print()
}
