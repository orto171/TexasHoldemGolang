package models

import (
	"math/rand"
)

var (
	SharedCards []Card
	PlayerCards []Card
)

type Card struct {
	Shape string
	Value int
}

func IsCardAlreadyDrawn(c Card) bool {
	allDrawnCards := append(PlayerCards, SharedCards...)
	for _, v := range allDrawnCards {
		if c.Shape == v.Shape && c.Value == v.Value {
			return true
		}
	}
	return false
}

func DrawShapeAndValue() Card {
	shapes := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	shape := shapes[rand.Intn(len(shapes))]
	value := rand.Intn(13) + 2
	return Card{Shape: shape, Value: value}
}

func DrawSharedCard() Card {
	for {
		card := DrawShapeAndValue()
		if !IsCardAlreadyDrawn(card) {
			SharedCards = append(SharedCards, card)
			return card
		}
	}
}

func DrawPlayerCard() Card {
	for {
		card := DrawShapeAndValue()
		if !IsCardAlreadyDrawn(card) {
			PlayerCards = append(PlayerCards, card)
			return card
		}
	}
}
