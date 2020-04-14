package dto

import "git.seevo.online/jgerlach/Games.Cards.API/model/memory"

type DrawCardsDto struct {
	Remaining	int				`json:"remaining"`
	Cards		[]memory.Card	`json:"cards"`
	Success		bool			`json:"success"`
	DeckId		string			`json:"deck_id"`
}
