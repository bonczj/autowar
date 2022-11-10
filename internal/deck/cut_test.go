package deck_test

import (
	"testing"

	"github.com/bonczj/autowar/internal/deck"
	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	d, err := deck.NewDeck()
	assert.NoError(t, err)
	assert.NotNil(t, d)
	assert.Equal(t, deck.DeckSize, d.Length())

	d.Cut()
	assert.NotNil(t, d)
	assert.Equal(t, deck.DeckSize, d.Length())
}
