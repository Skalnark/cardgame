package test

import (
	"fmt"
	"skalnark/cards/domain"
	database "skalnark/cards/providers"
	"skalnark/cards/service"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateTestDeckService() service.DeckService {
	DB, err := gorm.Open(sqlite.Open("../database/test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Couldn't open the database: ", err)
		return nil
	}
	database.DbMigrate(DB)

	return service.NewDeckService(DB)
}

func TestCreateDeck(t *testing.T) {
	deckService := CreateTestDeckService()

	input := []string{"2H", "AS", "JC", "QD", "10H"}
	indexes := []int{28, 40, 11, 25, 36}
	deck := &domain.Deck{}

	err := deckService.CreateDeck(deck, input)
	if err != nil {
		t.Errorf("got err %s, expected nil", err)
	}

	for i, c := range deck.Cards {
		if c.Index != indexes[i] {
			t.Errorf("got card.Index %d for %s, expected %d", c.Index, c.Code, indexes[i])
		}
	}

	if !deck.Shuffled {
		t.Errorf("got deck shuffled %t, expected true", deck.Shuffled)
	}

	if deck.Remaining != 5 {
		t.Errorf("got deck remmaining %d, expected 5", deck.Remaining)
	}

	input = []string{"QD"}
	deck = &domain.Deck{}
	err = deckService.CreateDeck(deck, input)
	if err != nil {
		t.Errorf("got err %s, expected nil", err)
	}

	deck = &domain.Deck{}
	err = deckService.CreateDeck(deck, nil)
	if err != nil {
		t.Errorf("got err %s, expected nil", err)
	}

	if deck.Remaining != 52 {
		t.Errorf("got %d cards for a default deck, expected 52", deck.Remaining)
	}

	if deck.Shuffled {
		t.Errorf("got deck shuffled %t, expected false", deck.Shuffled)
	}

	if deck.Remaining != 52 {
		t.Errorf("got deck remmaining %d, expected 52", deck.Remaining)
	}
}

func TestDrawCard(t *testing.T) {
	deckService := CreateTestDeckService()
	input := []string{"2C", "AS", "JD", "QH", "10S"}

	deck := &domain.Deck{}
	err := deckService.CreateDeck(deck, input)
	if err != nil {
		t.Errorf("got err %s, expected nil", err)
	}

	amount := 3
	remaining := deck.Remaining
	deckService.DrawCards(deck, amount)
	if err != nil {
		t.Errorf("got err %s, expected nil", err)
	}

	if deck.Shuffled {
		t.Errorf("got deck shuffled %t, expected false", deck.Shuffled)
	}

	if deck.Remaining != remaining-amount {
		t.Errorf("got deck remmaining %d, expected %d", deck.Remaining, remaining-amount)
	}
}

func TestOpenDeck(t *testing.T) {
	deckService := CreateTestDeckService()
	input := []string{"2H", "AS", "JC", "QD", "10H"}
	indexes := []int{28, 40, 11, 25, 36}
	deck := &domain.Deck{}

	err := deckService.CreateDeck(deck, input)
	if err != nil {
		t.Errorf("got err %s, expected nil", err)
	}

	id := deck.ID
	deck = &domain.Deck{ID: id}

	deckService.OpenDeck(deck)

	for i, c := range deck.Cards {
		if c.Index != indexes[i] {
			t.Errorf("got card.Index %d for %s, expected %d", c.Index, c.Code, indexes[i])
		}
	}

	if !deck.Shuffled {
		t.Errorf("got deck shuffled %t, expected true", deck.Shuffled)
	}

	if deck.Remaining != 5 {
		t.Errorf("got deck remmaining %d, expected 5", deck.Remaining)
	}
}
