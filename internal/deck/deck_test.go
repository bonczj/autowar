package deck_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bonczj/autowar/internal/cards"
	"github.com/bonczj/autowar/internal/deck"
)

func TestNewDeck(t *testing.T) {
	deck, err := deck.NewDeck()
	assert.NoError(t, err)
	assert.Equal(t, 52, deck.Length())

	// deck has not been randomized, so test first and last card
	assert.Equal(t, "two of "+string(cards.ClubsStr), deck.CardAt(0).String())
	assert.Equal(t, "ace of "+string(cards.SpadesStr), deck.CardAt(51).String())

	// test the two of 'suit' is every 13th card
	assert.Equal(t, "two of "+string(cards.DiamondsStr), deck.CardAt(13).String())
	assert.Equal(t, "two of "+string(cards.HeartsStr), deck.CardAt(26).String())
	assert.Equal(t, "two of "+string(cards.SpadesStr), deck.CardAt(39).String())
}

func TestShuffle(t *testing.T) {
	d, err := deck.NewDeck()
	assert.NoError(t, err)
	assert.NotNil(t, d)
	assert.Equal(t, deck.DeckSize, d.Length())

	d.Shuffle()
	assert.NotNil(t, d)
	assert.Equal(t, deck.DeckSize, d.Length())
}

func TestCut(t *testing.T) {
	d, err := deck.NewDeck()
	assert.NoError(t, err)
	assert.NotNil(t, d)
	assert.Equal(t, deck.DeckSize, d.Length())

	d.Cut()
	assert.NotNil(t, d)
	assert.Equal(t, deck.DeckSize, d.Length())
}
