package main

import "fmt"

type deck []string

func NewDeck() deck {

	card := []string{}
	cardsType := []string{"club", "hearts", "spades", "diamond"}
	cardsValues := []string{"one", "two", "three", "four"}

	for _, cardtype := range cardsType {
		for _, cardsValues := range cardsValues {
			card = append(card, cardsValues+" of "+cardtype)
		}
	}

	return card
}

func (d deck) print() {
	for _, value := range d {
		fmt.Println(value)
	}
}
