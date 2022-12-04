package models

import "strconv"

func GetCardValueSpecialCases(c Card) string {
	cardValuesToString := map[string]int{"jack": 11, "queen": 12, "king": 13, "ace": 15}
	for k, v := range cardValuesToString {
		if v == c.Value {
			return k
		}
	}
	return strconv.FormatInt(int64(c.Value), 10)
}

func ConvertCardToString(c Card) string {
	cardValueString := GetCardValueSpecialCases(c)
	return "Card shape: " + c.Shape + ", Card value: " + cardValueString
}

func ConvertCardsListToString(cards []Card) string {
	var cardsConvertedToString string
	for _, c := range cards {
		cardsConvertedToString = cardsConvertedToString +
			"\n" + ConvertCardToString(c)
	}
	return cardsConvertedToString
}
