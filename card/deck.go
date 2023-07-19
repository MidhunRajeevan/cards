package card

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func NewDeck() deck {

	// var card deck
	card := deck{}
	// card := []string{}
	cardsType := []string{"club", "hearts", "spades", "diamond"}
	cardsValues := []string{"one", "two", "three", "four"}

	for _, cardtype := range cardsType {
		for _, cardsValues := range cardsValues {
			card = append(card, cardsValues+" of "+cardtype)
		}
	}

	return card
}

func (d deck) Print() {
	for _, value := range d {
		fmt.Println(value)
	}
}

func (d deck) SaveToFile(fileName string) {
	s := strings.Join(d, ",")
	ioutil.WriteFile(fileName, []byte(s), 0644)
}

func (d deck) PickOne() string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(d))
	selectedCard := d[i]
	return selectedCard

}
