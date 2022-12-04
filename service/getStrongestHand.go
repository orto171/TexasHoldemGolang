package service

import (
	"Poker/models"
	"errors"
	"sort"
	"strconv"
)

func GetStrongestHand() (string, error) {
	error := ValidateCardsNumber()
	if error != nil {
		return "", error
	}
	cards := append(models.SharedCards, models.PlayerCards...)
	return getStrongestHandHandler(cards)

}

func getStrongestHandHandler(cards []models.Card) (string, error) {
	if isRoyalFlush(cards) {
		return "royal flush", nil
	} else if isStraightFlush(cards) {
		return "straight flush", nil
	} else if isFourOfKind(cards) {
		return "four of a kind", nil
	} else if isFlush, _ := isFlush(cards); isFlush {
		return "flush", nil
	} else if isStraight(cards) {
		return "straight", nil
	} else if threeOfKind, _ := isThreeOfKind(cards); threeOfKind {
		return "three of a kind", nil
	} else if isTwoPair(cards) {
		return "two pair", nil
	} else if isPair(cards) {
		return "pair", nil
	} else {
		return "high card", nil
	}
}

func ValidateCardsNumber() error {
	if len(models.SharedCards) != 5 {
		return errors.New("shared cards were not drawn")
	}
	if len(models.PlayerCards) != 2 {
		return errors.New("player cards were not drawn")
	}
	return nil
}

func isPair(cards []models.Card) bool {
	cardsValueCount := map[string]int{}
	for _, c := range cards {
		cardAsString := strconv.FormatInt(int64(c.Value), 10)
		if _, ok := cardsValueCount[cardAsString]; ok {
			return true
		} else {
			cardsValueCount[cardAsString] = 1
		}
	}
	return false
}

func isTwoPair(cards []models.Card) bool {
	cardsValueCount := map[string]int{}
	pairCount := 0
	for _, c := range cards {
		cardAsString := strconv.FormatInt(int64(c.Value), 10)
		if val, ok := cardsValueCount[cardAsString]; ok {
			if val == 1 {
				if pairCount == 1 {
					return true
				} else {
					pairCount++
					cardsValueCount[cardAsString]++
				}
			} else {
				cardsValueCount[cardAsString]++
			}
		} else {
			cardsValueCount[cardAsString] = 1
		}
	}
	return false
}

func isThreeOfKind(cards []models.Card) (bool, int) {
	cardsValueCount := map[string]int{}
	for _, c := range cards {
		cardAsString := strconv.FormatInt(int64(c.Value), 10)
		if val, ok := cardsValueCount[cardAsString]; ok {
			if val == 3 {
				return true, c.Value
			} else {
				cardsValueCount[cardAsString]++
			}
		} else {
			cardsValueCount[cardAsString] = 1
		}
	}
	return false, 0
}

func isStraight(cards []models.Card) bool {
	cardsNumbersValues := []int{}
	for _, c := range cards {
		cardsNumbersValues = append(cardsNumbersValues, c.Value)
	}
	sort.Ints(cardsNumbersValues)
	if checkStraightForSlice(cardsNumbersValues[0:5]) ||
		checkStraightForSlice(cardsNumbersValues[1:6]) ||
		checkStraightForSlice(cardsNumbersValues[2:]) {
		return true
	}
	return false
}

func checkStraightForSlice(cards []int) bool {
	for i := 0; i < len(cards)-1; i++ {
		if cards[i+1]-cards[i] > 1 {
			return false
		}
	}
	return true
}

func isFlush(cards []models.Card) (bool, string) {
	cardsShapesCount := map[string]int{"Spades": 0, "Clubs": 0, "Diamonds": 0, "Hearts": 0}
	for _, c := range cards {
		if val, ok := cardsShapesCount[c.Shape]; ok {
			if val == 5 {
				return true, c.Shape
			} else {
				cardsShapesCount[c.Shape]++
			}
		}
	}
	return false, ""
}

func isFullHouse(cards []models.Card) bool {
	if isThreeOfKind, threeTimesValue := isThreeOfKind(cards); isThreeOfKind {
		remainingCards := []models.Card{}
		for _, c := range cards {
			if c.Value == threeTimesValue {
				continue
			}
			remainingCards = append(remainingCards, c)
		}
		if isPair(remainingCards) {
			return true
		}
	}
	return false
}

func isFourOfKind(cards []models.Card) bool {
	cardsValueCount := map[string]int{}
	for _, c := range cards {
		cardAsString := strconv.FormatInt(int64(c.Value), 10)
		if val, ok := cardsValueCount[cardAsString]; ok {
			if val == 4 {
				return true
			} else {
				cardsValueCount[cardAsString]++
			}
		} else {
			cardsValueCount[cardAsString] = 1
		}
	}
	return false
}

func isStraightFlush(cards []models.Card) bool {
	isFlush, _ := isFlush(cards)
	if isFlush && isStraight(cards) {
		return true
	}
	return false
}

func isRoyalFlush(cards []models.Card) bool {
	isFlush, shape := isFlush(cards)
	if !isFlush {
		return false
	}
	onlyChosenShapeCards := excludeOtherShapesCards(cards, shape)
	if checkIfNumberExistInList(onlyChosenShapeCards, 10) &&
		checkIfNumberExistInList(onlyChosenShapeCards, 11) &&
		checkIfNumberExistInList(onlyChosenShapeCards, 12) &&
		checkIfNumberExistInList(onlyChosenShapeCards, 13) &&
		checkIfNumberExistInList(onlyChosenShapeCards, 14) {
		return true
	} else {
		return false
	}
}

func checkIfNumberExistInList(cardValues []int, number int) bool {
	for _, value := range cardValues {
		if value == number {
			return true
		}
	}
	return false
}

func excludeOtherShapesCards(cards []models.Card, shape string) []int {
	onlyChosenShapeCards := []int{}
	for _, c := range cards {
		if c.Shape == shape {
			onlyChosenShapeCards = append(onlyChosenShapeCards, c.Value)
		}
	}
	return onlyChosenShapeCards
}
