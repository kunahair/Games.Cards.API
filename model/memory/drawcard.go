package memory

type DrawCard struct {
	Remaining	int		`json:"remaining"`
	Cards		[]Card	`json:"cards"`
	Success		bool	`json:"success"`
	DeckId		string	`json:"deck_id"`
}
