package cards

import (
	"fmt"
)

type Value int

const (
	Two = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Value Value
	Title string
}

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades

	ClubsStr    string = "♣"
	DiamondsStr string = "♦"
	HeartsStr   string = "♥"
	SpadesStr   string = "♠"
)

func NewCard(value Value, suit Suit) (Card, error) {
	if value < Two || value > Ace {
		return Card{}, fmt.Errorf("value must be between two and ace: %d", value)
	}
	if suit < Clubs || suit > Spades {
		return Card{}, fmt.Errorf("suit must be club, diamond, heart or spade: %d", suit)
	}

	v := valueAsString(value)

	return Card{
		Value: value,
		Title: v + " " + suitToString(suit),
	}, nil
}

func valueAsString(value Value) string {
	str := ""

	switch value {
	case Two:
		str = "two of"
	case Three:
		str = "three of"
	case Four:
		str = "four of"
	case Five:
		str = "five of"
	case Six:
		str = "six of"
	case Seven:
		str = "seven of"
	case Eight:
		str = "eight of"
	case Nine:
		str = "nine of"
	case Ten:
		str = "ten of"
	case Jack:
		str = "jack of"
	case Queen:
		str = "queen of"
	case King:
		str = "king of"
	case Ace:
		str = "ace of"
	}

	return str
}

func suitToString(suit Suit) string {
	str := ""

	switch suit {
	case Clubs:
		str = ClubsStr
	case Diamonds:
		str = DiamondsStr
	case Hearts:
		str = HeartsStr
	case Spades:
		str = SpadesStr
	}

	return str
}

func (c Card) String() string {
	return c.Title
}
