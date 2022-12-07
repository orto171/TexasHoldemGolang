package models

import (
	"math/rand"
	"time"
)

var (
	Shapes      = []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	SharedCards []Card
	PlayerCards []Card
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Shape string
	Value int
}

func CreateDeck() *Deck {
	deck := Deck{Cards: []Card{}}
	for _, shape := range Shapes {
		for i := 2; i <= 14; i++ {
			card := Card{Shape: shape, Value: i}
			deck.Cards = append(deck.Cards, card)
		}
	}
	return &deck
}

func (deck *Deck) DrawRandomCard() Card {
	rand.Seed(time.Now().Unix())
	drawnCardIndex := rand.Intn(len(deck.Cards))
	drawnCard := deck.Cards[drawnCardIndex]
	deck.Cards = append(deck.Cards[:drawnCardIndex], deck.Cards[drawnCardIndex+1:]...)
	return drawnCard
}
