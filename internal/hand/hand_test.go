package hand_test

import (
	"testing"

	"github.com/bonczj/autowar/internal/cards"
	"github.com/bonczj/autowar/internal/hand"
	"github.com/stretchr/testify/assert"
)

func TestNewHand(t *testing.T) {
	hand := hand.NewHand()
	assert.Equal(t, 0, hand.LenHand())
	assert.Equal(t, 0, hand.LenDiscard())
	assert.True(t, hand.Busted())
}

func TestAddingToHand(t *testing.T) {
	hand := hand.NewHand()
	hand.Add([]cards.Card{cards.Card{}})
	hand.Add([]cards.Card{cards.Card{}})
	hand.Add([]cards.Card{cards.Card{}})
	assert.Equal(t, 3, hand.LenHand())
	assert.Equal(t, 0, hand.LenDiscard())
	assert.False(t, hand.Busted())
}

func TestAddingToDiscard(t *testing.T) {
	hand := hand.NewHand()
	list := make([]cards.Card, 0, 10)
	for i := 0; i < 10; i++ {
		list = append(list, cards.Card{})
	}

	hand.AddToDiscard(list)
	assert.Equal(t, 0, hand.LenHand())
	assert.Equal(t, 10, hand.LenDiscard())
	assert.False(t, hand.Busted())

	// take the next card will move discard to the hand then take the top card
	card := hand.Draw()
	assert.NotNil(t, card)
	assert.Equal(t, 9, hand.LenHand())
	assert.Equal(t, 0, hand.LenDiscard())
	assert.False(t, hand.Busted())
}
