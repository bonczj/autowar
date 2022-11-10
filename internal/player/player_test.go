package player_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bonczj/autowar/internal/player"
)

func TestNewPlayer(t *testing.T) {
	p := player.NewPlayer("one")
	assert.NotEmpty(t, p.String())
	assert.Equal(t, 0, p.Hand.LenHand())
	assert.Equal(t, 0, p.Hand.LenDiscard())
	assert.True(t, p.Hand.Busted())
}
