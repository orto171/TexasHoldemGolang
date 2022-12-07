package service

import (
	"Poker/models"
)

func DrawSharedCards(deck *models.Deck) string {

	cards, cardsStringRepresntation := DrawCards(deck, 5)
	models.SharedCards = cards
	return cardsStringRepresntation
}

func DrawPlayerCards(deck *models.Deck) string {
	cards, cardsStringRepresntation := DrawCards(deck, 2)
	models.PlayerCards = cards
	return cardsStringRepresntation
}

func DrawCards(deck *models.Deck, numberOfCards int) ([]models.Card, string) {
	cards := []models.Card{}
	for i := 0; i < numberOfCards; i++ {
		card := deck.DrawRandomCard()
		cards = append(cards, card)
	}
	return cards, models.ConvertCardsListToString(cards)
}
