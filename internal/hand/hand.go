package hand

import (
	"crypto/sha256"
	"fmt"

	"github.com/bonczj/autowar/internal/cards"
)

type Hand struct {
	cards   []cards.Card
	discard []cards.Card
}

func NewHand() Hand {
	return Hand{
		cards:   make([]cards.Card, 0),
		discard: make([]cards.Card, 0),
	}
}

// Add adds a card to the top of the hand
func (h *Hand) Add(list []cards.Card) {
	h.cards = append(h.cards, list...)
}

// AddToDiscard adds a set of cards to the discard pile
func (h *Hand) AddToDiscard(list []cards.Card) {
	h.discard = append(h.discard, list...)
}

func (h Hand) LenHand() int {
	return len(h.cards)
}

func (h Hand) LenDiscard() int {
	return len(h.discard)
}

func (h Hand) Length() int {
	return h.LenHand() + h.LenDiscard()
}

func (h Hand) Busted() bool {
	return h.LenHand() == 0 && h.LenDiscard() == 0
}

// Draw returns the first card in the hand. If the players
// hand is empty, but the discard pile is not, move the discard
// into their hand and take the first card. If both the hand and
// discard pile are empty, return nil (no card).
func (h *Hand) Draw() *cards.Card {
	if h.Busted() {
		return nil
	} else if h.LenHand() == 0 {
		h.cards = h.discard
		h.discard = make([]cards.Card, 0)
	}

	card := h.cards[0]
	h.cards = h.cards[1:]

	return &card
}

// Hash creates a SHA265 hash of all the cards in the hand + discard.
func (h Hand) Hash() string {
	input := ""
	for _, c := range h.cards {
		input += c.String()
	}
	for _, c := range h.discard {
		input += c.String()
	}

	sha_256 := sha256.New()
	sha_256.Write([]byte(input))
	return fmt.Sprintf("%x", sha_256.Sum(nil))
}
