package memory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/tkanos/gonfig"
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	Shuffled	bool	`json:"shuffled"`
	Remaining	int		`json:"remaining"`
	Success		bool	`json:"success"`
	DeckId		string	`json:"deck_id"`

	RemainingCards	[]Card	`json:"remaining_cards"`
}

func NewDeck(url string) (Deck, error) {
	deck := Deck{
		RemainingCards: []Card{},
	}

	cards := CardsTopLevel{}
	err := gonfig.GetConf("./cards.json", &cards)
	if err != nil {
		return deck, err
	}

	for i := 0; i < len(cards.Cards); i++ {
		cards.Cards[i].Image = url +  "/static/img/" + cards.Cards[i].Code + ".png"
		cards.Cards[i].Images.Png = url +  "/static/img/" + cards.Cards[i].Code + ".png"
		cards.Cards[i].Images.Svg = url +  "/static/img/" + cards.Cards[i].Code + ".svg"
	}

	uuidWithHyphen := uuid.New()
	uuidString := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	deck.RemainingCards = cards.Cards
	deck.DeckId = uuidString
	deck.Remaining = len(cards.Cards)
	deck.Shuffled = true
	deck.Success = true

	return deck, nil
}

func (d *Deck) ShuffleCards() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.RemainingCards), func(i, j int) {
		d.RemainingCards[i], d.RemainingCards[j] = d.RemainingCards[j], d.RemainingCards[i]
	})
}

func (d *Deck) DrawCard(count int) ([]Card, error) {

	cards := []Card{}

	if count > len(d.RemainingCards) {
		return cards, errors.New("not enough cards remain for count")
	}

	if count < 1 {
		return cards, nil
	}

	for i := 0; i < count; i++ {
		card, deck := d.RemainingCards[0], d.RemainingCards[1:]
		d.RemainingCards = deck
		cards = append(cards, card)
	}

	d.Remaining = len(d.RemainingCards)

	return cards, nil
}
