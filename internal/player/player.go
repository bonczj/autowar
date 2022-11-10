package player

import (
	"fmt"

	"github.com/bonczj/autowar/internal/hand"
)

type Player struct {
	name   string
	Hand   hand.Hand
	hashes map[string]bool
}

func NewPlayer(name string) Player {
	return Player{
		name:   name,
		Hand:   hand.NewHand(),
		hashes: make(map[string]bool),
	}
}

func (p Player) String() string {
	return fmt.Sprintf("Player %s has %d cards left with a discard pile of %d cards",
		p.name, p.Hand.LenHand(), p.Hand.LenDiscard())
}

// Hash creates a hash of the players current hand
func (p *Player) Hash() {
	hash := p.Hand.Hash()
	p.hashes[hash] = true
}

// SeenHand checks if the current hand has been seen by the player before
func (p Player) SeenHand() bool {
	hash := p.Hand.Hash()
	_, ok := p.hashes[hash]
	return ok
}
