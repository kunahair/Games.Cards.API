package controller

import (
	"git.seevo.online/jgerlach/Games.Cards.API/config"
	"git.seevo.online/jgerlach/Games.Cards.API/model/dto"
	"git.seevo.online/jgerlach/Games.Cards.API/model/memory"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func DrawCardFromDeck(config config.Configuration, engine *memory.DeckEngine) gin.HandlerFunc {
	return func(c *gin.Context) {

		deckId := c.Param("deckId")

		if len(deckId) < 1 || deckId == "" {
			JsonMessageWithStatus(c, http.StatusBadRequest, "deck id invalid")
			return
		}

		countString := c.DefaultQuery("count", "0")
		count, err := strconv.Atoi(countString)
		if err != nil {
			JsonMessageWithStatus(c, http.StatusBadRequest, "invalid count for card draw")
			return
		}

		var deck memory.Deck
		deckFound := false
		for _, val := range engine.Decks {
			if val.DeckId == deckId {
				deckFound = true
				deck = val
			}
		}

		if !deckFound {
			JsonMessageWithStatus(c, http.StatusNotFound, "deck not found")
			return
		}

		if len(deck.RemainingCards) < count {
			JsonMessageWithStatus(c, http.StatusBadRequest, "not enough cards remain for draw count")
			return
		}

		cards, err := deck.DrawCard(count)
		if err != nil {
			JsonMessageWithStatus(c, http.StatusInternalServerError, "unable to draw card")
			return
		}

		engine.UpdateDeck(deck)

		cardsDto := dto.DrawCardsDto{
			Remaining: deck.Remaining,
			Cards: 	   cards,
			Success:   deck.Success,
			DeckId:    deck.DeckId,
		}

		c.JSON(http.StatusOK, cardsDto)

	}
}



func NewDeck(config config.Configuration, engine *memory.DeckEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create new Deck
		deck, err := memory.NewDeck(config.ApiUrl)
		if err != nil {
			log.Fatal(err)
		}

		deck.ShuffleCards()

		engine.AddDeck(deck)

		deckDto := dto.DeckDto{
			Shuffled:  deck.Shuffled,
			Remaining: deck.Remaining,
			Success:   deck.Success,
			DeckId:    deck.DeckId,
		}

		c.JSON(http.StatusCreated, deckDto)
	}
}
