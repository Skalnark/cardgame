package database

import (
	"skalnark/cards/domain"

	"gorm.io/gorm"
)

type DeckRepository interface {
	NewDeck(*domain.Deck, []domain.Card)
	GetDeckById(*domain.Deck)
	DrawCards(*domain.Deck, int)
}

func NewDeckRepository(db *gorm.DB) DeckRepository {
	return &Repository{db: db}
}

func (r *Repository) NewDeck(deck *domain.Deck, cards []domain.Card) {

	for i, j := 0, 1; j < len(cards); i, j = i+1, j+1 {
		if cards[i].Index > cards[j].Index {
			deck.Shuffled = true
			break
		}
	}

	deck.Remaining = len(cards)
	r.db.Create(deck)

	for _, c := range cards {
		c.DeckID = deck.ID
		r.db.Create(&c)
	}

	r.db.Find(&deck.Cards, domain.Card{DeckID: deck.ID})
}

func (r *Repository) GetDeckById(deck *domain.Deck) {

	r.db.Find(deck, domain.Deck{ID: deck.ID})
	r.db.Find(&deck.Cards, domain.Card{DeckID: deck.ID})
}

func (r *Repository) DrawCards(deck *domain.Deck, count int) {

	var drawCards []domain.Card
	r.db.Find(&drawCards, domain.Card{DeckID: deck.ID})

	for i, j := 0, 0; i < count && j < len(drawCards); i++ {
		r.db.Where("deck_id = ?", deck.ID).Delete(&drawCards[i])
		j++
	}

	var aux []domain.Card

	r.db.Find(&aux, domain.Card{DeckID: deck.ID})

	deck.Shuffled = false
	for i, j := 0, 1; j < len(aux); i, j = i+1, j+1 {
		if aux[i].Index > aux[j].Index {
			deck.Shuffled = true
			break
		}
	}

	deck.Remaining = len(aux)

	deck.Cards = []domain.Card{}
	r.db.Model(&deck).Omit("cards").Updates(
		domain.Deck{Remaining: deck.Remaining, Shuffled: deck.Shuffled})
	deck.Cards = aux
}
