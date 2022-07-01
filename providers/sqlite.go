package database

import (
	"skalnark/cards/domain"

	"gorm.io/gorm"
)

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&domain.Card{})
	db.AutoMigrate(&domain.Deck{})
}
