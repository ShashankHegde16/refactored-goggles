package main

func main() {

	cards := addDeck()
	cards.shuffle()
	cards.print()
}

func newName() string {
	return "Shashank"
}
