package deck

import "github.com/bonczj/autowar/internal/cards"

// Deal will shuffle/cut the deck a few times, then
// distribute the cards evenly between the two players.
func (d *Deck) Deal() ([]cards.Card, []cards.Card) {
	hand1 := make([]cards.Card, 0, d.Length()/2)
	hand2 := make([]cards.Card, 0, d.Length()/2)

	for i := 0; i < 3; i++ {
		d.Shuffle()
	}

	for i, c := range d.cards {
		if i%2 == 0 {
			hand1 = append(hand1, c)
		} else {
			hand2 = append(hand2, c)
		}
	}

	// remove the internal cards as the players will now have them
	d.cards = make([]cards.Card, 0)

	return hand1, hand2
}
