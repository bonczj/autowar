package cards_test

import (
	"testing"

	"github.com/bonczj/autowar/internal/cards"
	"github.com/stretchr/testify/assert"
)

func TestTwoOfSpades(t *testing.T) {
	card, err := cards.NewCard(cards.Two, cards.Spades)
	assert.NoError(t, err)
	assert.Equal(t, "two of "+string(cards.SpadesStr), card.String())
}

func TestAceOfHearts(t *testing.T) {
	card, err := cards.NewCard(cards.Ace, cards.Hearts)
	assert.NoError(t, err)
	assert.Equal(t, "ace of "+string(cards.HeartsStr), card.String())
}
