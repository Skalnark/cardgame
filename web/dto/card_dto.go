package dto

import "skalnark/cards/domain"

type DeckDto struct {
	Id        string    `json:"id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     []CardDto `json:"cards"`
}

type CardDto struct {
	Value string `json:"value"`
	Suite string `json:"suite"`
	Code  string `json:"code"`
}

func CreateResponse(deck *domain.Deck) (response DeckDto) {
	response = DeckDto{
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
		Id:        deck.ID,
	}
	for _, c := range deck.Cards {
		card := CardDto{Value: c.Value, Suite: c.Suite, Code: c.Code}
		response.Cards = append(response.Cards, card)
	}
	return response
}
