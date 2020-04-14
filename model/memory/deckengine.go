package memory

import "sync"

type DeckEngine struct {
	Decks	[]Deck

	sync.Mutex
}

func (de *DeckEngine) AddDeck(deck Deck) {
	de.Lock()
	defer de.Unlock()

	de.Decks = append(de.Decks, deck)
}

func (de *DeckEngine) UpdateDeck(deck Deck) {
	de.Lock()
	defer de.Unlock()

	for i := 0; i < len(de.Decks); i++ {
		if de.Decks[i].DeckId == deck.DeckId {
			de.Decks[i] = deck
			break
		}
	}

}
