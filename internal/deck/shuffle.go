package deck

import (
	"math/rand"
	"time"

	"github.com/bonczj/autowar/internal/cards"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixMicro()))
}

// Shuffle simulates a person shuffling cards.
// Split the deck into roughly half.
// Combine cards from the top of each 'half' into the final deck.
// Randomly select from either side and for a small amount of cards
// at a time.
func (d *Deck) Shuffle() {
	split := len(d.cards)/2 - random.Intn(10)
	left := d.cards[:split]
	right := d.cards[split:]
	shuffled := make([]cards.Card, 0, len(d.cards))
	d.Cut()

	for {
		switch {
		case len(left) == 0 && len(right) == 0:
			d.cards = shuffled
			return
		case len(left) > 0 && len(right) > 0:
			// how many cards are we going to drop?
			count := 1 + random.Intn(4)

			// randomly, pick left or right
			if random.Intn(2) == 0 {
				// go left
				if count < len(left) {
					shuffled = append(shuffled, left[:count]...)
					left = left[count:]
				} else {
					shuffled = append(shuffled, left...)
					left = make([]cards.Card, 0)
				}
			} else {
				// go right
				if count < len(right) {
					shuffled = append(shuffled, right[:count]...)
					right = right[count:]
				} else {
					shuffled = append(shuffled, right...)
					right = make([]cards.Card, 0)
				}
			}
		case len(left) > 0 && len(right) == 0:
			shuffled = append(shuffled, left...)
			left = make([]cards.Card, 0)
		case len(right) > 0 && len(left) == 0:
			shuffled = append(shuffled, right...)
			right = make([]cards.Card, 0)
		}
	}
}
