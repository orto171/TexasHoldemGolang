package service

import (
	"Poker/models"
)

func DrawSharedCards() string {
	models.SharedCards = []models.Card{}
	for i := 0; i < 5; i++ {
		models.DrawSharedCard()
	}
	cardsConvertedToString := models.ConvertCardsListToString(models.SharedCards)
	return cardsConvertedToString
}

func DrawPlayerCards() string {
	models.PlayerCards = []models.Card{}
	for i := 0; i < 2; i++ {
		models.DrawPlayerCard()
	}
	cardsConvertedToString := models.ConvertCardsListToString(models.PlayerCards)
	return cardsConvertedToString
}
