// package war is the game logic to start and play a game
// of war until either one player wins or they have a tie.
// A tie is defined as in a 'war' that starts another 'war',
// but one of the players runs out of cards.
package war

import (
	"log"
	"math"

	"github.com/bonczj/autowar/internal/cards"
	"github.com/bonczj/autowar/internal/deck"
	"github.com/bonczj/autowar/internal/player"
)

type War struct {
	PlayerOne player.Player
	PlayerTwo player.Player
	Rounds    int
}

type Winner int

const (
	Loop Winner = iota - 1
	WinnerTie
	WinnerPlayerOne
	WinnerPlayerTwo
	Continue
	Error
)

type Result struct {
	Win    Winner `csv:"winner"`
	Rounds int    `csv:"rounds"`
}

// NewWar will create a new deck, shuffle the deck then deal
// it between the players. The resulting war is ready to play
// with two players who each have half of the deck in their hand.
func NewWar() (*War, error) {
	playerOne := player.NewPlayer("one")
	playerTwo := player.NewPlayer("two")
	deck, err := deck.NewDeck()
	if err != nil {
		return nil, err
	}

	cards1, cards2 := deck.Deal()
	playerOne.Hand.Add(cards1)
	playerTwo.Hand.Add(cards2)

	return &War{
		PlayerOne: playerOne,
		PlayerTwo: playerTwo,
		Rounds:    0,
	}, nil
}

// Play runs the rules for a game of war. Apply these rules until
// either one player has no cards or a player cannot play due to
// not having enough cards.
//
// - Each player takes the top card
// - Player with the highest value card takes the cards
// - If the cards are the same value,
// - - each player draws 3 new cards
// - - each player draws another card and compare
// - - winner takes *all* cards from the war
// - - if the new cards are tied, run the war again until a winner
// - - if a player runs out of cards after the comparison in a war, the game is a tie
// - - if a player does not have enough cards for a war, use all of their cards for a short war
func (w *War) Play() Winner {
	for {
		if w.Rounds%1000 == 0 {
			log.Printf("Round %d", w.Rounds)
			log.Printf("Player 1 has %d cards", w.PlayerOne.Hand.Length())
			log.Printf("Player 2 has %d cards", w.PlayerTwo.Hand.Length())
			log.Println()
		}
		if w.PlayerOne.Hand.Busted() {
			return WinnerPlayerTwo
		}
		if w.PlayerTwo.Hand.Busted() {
			return WinnerPlayerOne
		}

		cards1 := make([]cards.Card, 0)
		cards2 := make([]cards.Card, 0)

		result := w.playRound(cards1, cards2, false)
		if result != Continue {
			return result
		}

		// if we have seen *both* hands before, just bail as we are likely going to loop
		if w.PlayerOne.SeenHand() && w.PlayerTwo.SeenHand() {
			return Loop
		}

		w.PlayerOne.Hash()
		w.PlayerTwo.Hash()
	}
}

func (w *War) playRound(cards1 []cards.Card, cards2 []cards.Card, inWar bool) Winner {
	if !inWar {
		w.Rounds++
	}

	// if we are in a war and one of the players has run out of cards, then we call it a tie
	if (w.PlayerOne.Hand.Busted() || w.PlayerTwo.Hand.Busted()) && inWar {
		return WinnerTie
	}

	// draw cards for both players. The checks above will prevent the draw from returning nil.
	cards1 = append(cards1, *w.PlayerOne.Hand.Draw())
	cards2 = append(cards2, *w.PlayerTwo.Hand.Draw())
	value1 := cards1[len(cards1)-1].Value
	value2 := cards2[len(cards2)-1].Value

	switch {
	case value1 == value2:
		warSize := 3

		// take up to 3 more cards as long as both players have at least 4 cards.
		// otherwise, the N-1 where N is the smallest hand left
		// if no cards are left, it's a tie
		if w.PlayerOne.Hand.Length() < 4 || w.PlayerTwo.Hand.Length() < 4 {
			warSize = int(math.Min(float64(w.PlayerOne.Hand.Length()), float64(w.PlayerTwo.Hand.Length()))) - 1
		}

		for i := 0; i < warSize; i++ {
			cards1 = append(cards1, *w.PlayerOne.Hand.Draw())
			cards2 = append(cards2, *w.PlayerTwo.Hand.Draw())
		}

		return w.playRound(cards1, cards2, true)
	case value1 > value2:
		w.PlayerOne.Hand.Add(cards1)
		w.PlayerOne.Hand.Add(cards2)
		return Continue
	case value1 < value2:
		w.PlayerTwo.Hand.Add(cards1)
		w.PlayerTwo.Hand.Add(cards2)
		return Continue
	}

	// should never hit this branch, but the switch does not know that
	log.Printf("Tie with %d cards outstanding", len(cards1)+len(cards2))
	return WinnerTie
}
