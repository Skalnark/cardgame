package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"skalnark/cards/domain"
	"skalnark/cards/web/dto"
	"strconv"
	"strings"
)

func (web *web) CreateDeck(w http.ResponseWriter, r *http.Request) {
	cardParameter := r.URL.Query().Get("cards")
	cards := strings.Split(cardParameter, ",")
	var response []byte

	deck := &domain.Deck{}

	err := web.service.CreateDeck(deck, cards)
	if err != nil {
		response = dto.NewError("error creating the deck: "+err.Error(),
			http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	responseObj := dto.CreateResponse(deck)

	response, err = json.Marshal(responseObj)
	if err != nil {
		response = dto.NewError("internal server error: "+err.Error(),
			http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (web *web) OpenDeck(w http.ResponseWriter, r *http.Request) {

	deckId := r.URL.Query().Get("deck_id")
	var response []byte

	if deckId == "" {
		response = dto.NewError("missing parameters: deck_id",
			http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	deck := &domain.Deck{ID: deckId}

	web.service.OpenDeck(deck)

	responseObj := dto.CreateResponse(deck)

	response, err := json.Marshal(responseObj)
	if err != nil {
		response = dto.NewError("internal server error: "+err.Error(),
			http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (web *web) DrawCards(w http.ResponseWriter, r *http.Request) {
	var response []byte

	countParameter := r.URL.Query().Get("count")
	if countParameter == "" {
		response = dto.NewError("missing parameters: count", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	count, err := strconv.Atoi(countParameter)
	if err != nil {
		response = dto.NewError("invalid parameters: count", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	deckId := r.URL.Query().Get("deck_id")
	if deckId == "" {
		response = dto.NewError("missing parameters: deck_id", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	deck := &domain.Deck{ID: deckId}

	web.service.DrawCards(deck, count)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	responseObj := dto.CreateResponse(deck)

	response, err = json.Marshal(responseObj)
	if err != nil {
		response = dto.NewError(err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
