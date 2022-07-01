package service

import (
	"errors"
	"skalnark/cards/domain"
	sqlite "skalnark/cards/providers"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type DeckService interface {
	CreateDeck(*domain.Deck, []string) error
	DrawCards(*domain.Deck, int)
	OpenDeck(*domain.Deck)
}

type deckService struct {
	repo sqlite.DeckRepository
}

func NewDeckService(db *gorm.DB) DeckService {
	return &deckService{
		repo: sqlite.NewDeckRepository(db),
	}
}

func (service *deckService) CreateDeck(deck *domain.Deck, r []string) (err error) {
	var cards []domain.Card

	for _, c := range r {
		card, err := parseCode(c)
		if err != nil {
			return err
		}

		cards = append(cards, card)
	}

	if len(cards) == 0 {
		DefaultDeck(&cards)
	}

	service.repo.NewDeck(deck, cards)
	return nil
}

func (service *deckService) DrawCards(deck *domain.Deck, count int) {
	service.repo.DrawCards(deck, count)
}

func (service *deckService) OpenDeck(deck *domain.Deck) {
	service.repo.GetDeckById(deck)
}

func parseCode(code string) (card domain.Card, err error) {

	length := len(code)
	if length <= 1 || length > 3 {
		err = errors.New("invalid card code: " + code)
		return card, err
	}

	var value string
	if length > 2 {
		value = code[0:2]
	} else {
		value = code[0:1]
	}

	iValue, err := strconv.Atoi(value)
	if err != nil {
		err = nil
		switch strings.ToUpper(value) {
		case "A":
			value = "ACE"
		case "J":
			value = "JACK"
		case "Q":
			value = "QUEEN"
		case "K":
			value = "KING"
		default:
			err = errors.New("invalid card value: " + value)
		}
	}
	if err != nil {
		return card, err
	}

	if iValue < 2 && iValue > 10 {
		err = errors.New("invalid card value: " + value)
		return card, err
	}

	var suite string

	switch code[length-1:] {
	case "C":
		suite = "CLUBS"
	case "D":
		suite = "DIAMONDS"
	case "H":
		suite = "HEARTS"
	case "S":
		suite = "SPADES"
	default:
		err = errors.New("invalid suite value: " + code[length-1:])
	}

	if err != nil {
		return card, err
	}

	index := getIndex(value, suite)

	card = domain.Card{Value: value, Suite: suite, Code: code, Index: index}

	return card, err
}

func getIndex(value string, suite string) int {
	var suiteAdd int

	switch suite {
	case "CLUBS":
		suiteAdd = 0
	case "DIAMONDS":
		suiteAdd = 13
	case "HEARTS":
		suiteAdd = 26
	case "SPADES":
		suiteAdd = 39
	}

	cardAdd, err := strconv.Atoi(value)
	if err != nil {
		switch value {
		case "ACE":
			cardAdd = 1
		case "JACK":
			cardAdd = 11
		case "QUEEN":
			cardAdd = 12
		case "KING":
			cardAdd = 13
		}
	}
	return cardAdd + suiteAdd
}

func DefaultDeck(cards *[]domain.Card) {

	var value string
	var suite string

	for i := 0; i < 4; i++ {
		for j := 1; j <= 13; j++ {
			if j == 1 {
				value = "ACE"
			} else if j > 10 {
				if j%11 == 0 {
					value = "JACK"
				} else if j%12 == 0 {
					value = "QUEEN"
				} else if j%13 == 0 {
					value = "KING"
				}
			} else {
				value = strconv.Itoa(j)
			}

			switch i {
			case 0:
				suite = "CLUBS"
			case 1:
				suite = "DIAMONDS"
			case 2:
				suite = "HEARTS"
			case 3:
				suite = "SPADES"
			}
			code := genCode(value, suite)
			card := domain.Card{Value: value, Suite: suite, Code: code, Index: i}
			*cards = append(*cards, card)
		}
	}
}

func genCode(value string, suite string) string {
	len := len(value)
	if len > 2 {
		return value[:1] + suite[:1]
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		return value[:1] + suite[:1]
	}

	return strconv.Itoa(v) + suite[:1]
}
