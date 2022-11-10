package war_test

import (
	"testing"

	"github.com/bonczj/autowar/internal/cards"

	"github.com/bonczj/autowar/internal/player"

	"github.com/bonczj/autowar/internal/deck"

	"github.com/bonczj/autowar/internal/war"
	"github.com/stretchr/testify/assert"
)

func TestNewWar(t *testing.T) {
	w, err := war.NewWar()
	assert.Nil(t, err)
	assert.NotNil(t, w)
	assert.Equal(t, deck.DeckSize/2, w.PlayerOne.Hand.LenHand())
	assert.Equal(t, 0, w.PlayerOne.Hand.LenDiscard())
	assert.False(t, w.PlayerOne.Hand.Busted())
	assert.Equal(t, deck.DeckSize/2, w.PlayerTwo.Hand.LenHand())
	assert.Equal(t, 0, w.PlayerTwo.Hand.LenDiscard())
	assert.False(t, w.PlayerTwo.Hand.Busted())
}

func TestSimpleWar(t *testing.T) {
	w := war.War{
		PlayerOne: player.NewPlayer("one"),
		PlayerTwo: player.NewPlayer("two"),
		Rounds:    0,
	}

	// player one wins every hand
	card1, _ := cards.NewCard(cards.Ace, cards.Hearts)
	card2, _ := cards.NewCard(cards.Two, cards.Diamonds)
	w.PlayerOne.Hand.Add([]cards.Card{card1})
	w.PlayerTwo.Hand.Add([]cards.Card{card2})
	// repeat same cards to ensure always winning
	w.PlayerOne.Hand.Add([]cards.Card{card1})
	w.PlayerTwo.Hand.Add([]cards.Card{card2})
	w.PlayerOne.Hand.Add([]cards.Card{card1})
	w.PlayerTwo.Hand.Add([]cards.Card{card2})

	winner := w.Play()
	assert.Equal(t, war.WinnerPlayerOne, winner)
	assert.Equal(t, 3, w.Rounds)
}

func TestPlayerTwoWinsWar(t *testing.T) {
	w := war.War{
		PlayerOne: player.NewPlayer("one"),
		PlayerTwo: player.NewPlayer("two"),
		Rounds:    0,
	}

	cards1 := make([]cards.Card, 0)
	cards2 := make([]cards.Card, 0)

	// first card tie to start war
	card1, _ := cards.NewCard(cards.Two, cards.Hearts)
	card2, _ := cards.NewCard(cards.Two, cards.Diamonds)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)

	// three cards where player 1 has higher value, but hidden due to war
	card1, _ = cards.NewCard(cards.Ten, cards.Hearts)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)

	// last card player 2 has higher value and takes all
	card2, _ = cards.NewCard(cards.Jack, cards.Hearts)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)

	// adding cards to
	w.PlayerOne.Hand.Add(cards1)
	w.PlayerTwo.Hand.Add(cards2)

	winner := w.Play()
	assert.Equal(t, war.WinnerPlayerTwo, winner)
	assert.Equal(t, 1, w.Rounds)
}

func TestTie(t *testing.T) {
	// start a war that spawns a second war. Player two has no more cards to continue the
	// war, so we end in a tie
	w := war.War{
		PlayerOne: player.NewPlayer("one"),
		PlayerTwo: player.NewPlayer("two"),
		Rounds:    0,
	}

	cards1 := make([]cards.Card, 0)
	cards2 := make([]cards.Card, 0)

	// first card tie to start war
	card1, _ := cards.NewCard(cards.Two, cards.Hearts)
	card2, _ := cards.NewCard(cards.Two, cards.Diamonds)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)

	// three cards where player 1 has higher value, but hidden due to war
	card1, _ = cards.NewCard(cards.Ten, cards.Hearts)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)

	// last card starts another war
	card2, _ = cards.NewCard(cards.Ten, cards.Hearts)
	cards1 = append(cards1, card1)
	cards2 = append(cards2, card2)

	// player 1 has another card still
	cards1 = append(cards1, card1)

	// adding cards to
	w.PlayerOne.Hand.Add(cards1)
	w.PlayerTwo.Hand.Add(cards2)

	winner := w.Play()
	assert.Equal(t, war.WinnerTie, winner)
	assert.Equal(t, 1, w.Rounds)
}
