package main

import (
	"Poker/models"
	"Poker/service"
	"fmt"
	"log"
	"net/http"
)

var (
	deck *models.Deck
)

func dealCardsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if r.URL.Path == "/deal_shared_cards" {
		sharedCards := service.DrawSharedCards(deck)
		fmt.Fprintf(w, sharedCards)

	} else if r.URL.Path == "/deal_player_cards" {
		playerCards := service.DrawPlayerCards(deck)
		fmt.Fprintf(w, playerCards)

	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func getStrongestHandHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if r.URL.Path == "/get_strongest_hand" {
		strongestHand, error := service.GetStrongestHand()
		if error == nil {
			fmt.Fprintf(w, "strongest hand is: "+strongestHand)
		} else {
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

}

func main() {

	deck = models.CreateDeck()

	http.HandleFunc("/deal_shared_cards", dealCardsHandler)

	http.HandleFunc("/deal_player_cards", dealCardsHandler)

	http.HandleFunc("/get_strongest_hand", getStrongestHandHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
