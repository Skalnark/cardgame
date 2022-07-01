package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	ID     string
	Suite  string
	Value  string
	Code   string
	Index  int
	DeckID string
	Deck   Deck
}

type Deck struct {
	gorm.Model
	ID        string
	Shuffled  bool
	Cards     []Card
	Remaining int // for performance
}

func (u *Card) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()

	return
}

func (u *Deck) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()

	return
}
