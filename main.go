package main

import (
	"log"
)

func main() {
	cards := addDeck()
	cards.print()

	err := cards.writeToFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

}

func newName() string {
	return "Shashank"
}
