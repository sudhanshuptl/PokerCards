package main

func main() {
	cards := newDeck()
	// cards.saveToFile("mycards")
	// cards := getDeckFromFile("mycards")
	cards.print()
}
