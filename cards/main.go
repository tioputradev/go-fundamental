package main

func main() {

	newCards := newDeckFromFile("my_cards")

	newCards.shuffle()
	newCards.print()
}
