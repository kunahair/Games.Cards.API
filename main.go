package main

import (
	"git.seevo.online/jgerlach/Games.Cards.API/config"
	"git.seevo.online/jgerlach/Games.Cards.API/controller"
	"git.seevo.online/jgerlach/Games.Cards.API/model/memory"
	"github.com/gin-gonic/gin"
	"github.com/tkanos/gonfig"
)

var build string
var deckEngine memory.DeckEngine

func main() {

	configuration := config.Configuration{}

	if len(build) > 0 && build == "production" {
		err := gonfig.GetConf("./config.production.json", &configuration)
		if err != nil {
			panic(err)
		}
	} else {
		err := gonfig.GetConf("./config.dev.json", &configuration)
		if err != nil {
			panic(err)
		}
	}

	deckEngine = memory.DeckEngine{
		Decks: []memory.Deck{},
	}

	r := gin.Default()

	apiRoutes := r.Group("/api")
	{
		deckRoutes := apiRoutes.Group("/deck")
		{
			deckRoutes.POST("/new/shuffle", controller.NewDeck(configuration, &deckEngine))

			deckRoutes.GET("/:deckId/draw", controller.DrawCardFromDeck(configuration, &deckEngine))
		}
	}

	//Static Images
	r.Static("/static/img", "./static/img")

	r.Run(configuration.Port)


}
