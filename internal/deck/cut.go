package deck

import (
	"github.com/bonczj/autowar/internal/cards"
)

// Cut will split the deck into a few segments and then
// add them back in a different order. We are going to
// make three cuts then add them back as 3-1-2.
func (d *Deck) Cut() {
	cut := make([]cards.Card, 0, len(d.cards))

	first := 10 + random.Intn(10)
	second := 10 + random.Intn(10)

	cut = append(cut, d.cards[first+second:]...)
	cut = append(cut, d.cards[:first]...)
	cut = append(cut, d.cards[first:first+second]...)

	d.cards = cut
}
