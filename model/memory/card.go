package memory

type Card struct {
	Value	string		`json:"value"`
	Code	string		`json:"code"`
	Suit	string		`json:"suit"`
	Image	string		`json:"image"`
	Images	CardImages	`json:"images"`
}

type CardImages struct {
	Png	string	`json:"png"`
	Svg	string	`json:"svg"`
}
