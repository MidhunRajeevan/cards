package main

import (
	"fmt"

	card "github.com/MidhunRajeevan/cards/card"
)

func main() {
	fmt.Println("hello")
	cards := card.NewDeck()
	cards.Print()
	// cards.SaveToFile("newCrad")
	selectedCard := cards.PickOne()
	fmt.Println("selectedCard : ", selectedCard)
}
