package deck

import "github.com/bonczj/autowar/internal/cards"

const DeckSize = 52
const SuitSize = 4
const CardsPerSuit = DeckSize / SuitSize

type Deck struct {
	cards []cards.Card
}

func NewDeck() (*Deck, error) {
	deck := Deck{
		cards: make([]cards.Card, 0, DeckSize),
	}
	suits := make([]cards.Suit, 0, SuitSize)

	suits = append(suits, cards.Clubs)
	suits = append(suits, cards.Diamonds)
	suits = append(suits, cards.Hearts)
	suits = append(suits, cards.Spades)

	for _, s := range suits {
		cards, err := buildSuit(s)
		if err != nil {
			return nil, err
		}

		deck.cards = append(deck.cards, cards...)
	}

	return &deck, nil
}

func (d *Deck) Length() int {
	return len(d.cards)
}

func (d *Deck) CardAt(position int) *cards.Card {
	if position < 0 || position > d.Length() {
		return nil
	}

	return &cards.Card{
		Value: d.cards[position].Value,
		Title: d.cards[position].Title,
	}
}

func buildSuit(suit cards.Suit) ([]cards.Card, error) {
	deck := make([]cards.Card, 0, CardsPerSuit)
	values := make([]cards.Value, 0, CardsPerSuit)

	values = append(values, cards.Two)
	values = append(values, cards.Three)
	values = append(values, cards.Four)
	values = append(values, cards.Five)
	values = append(values, cards.Six)
	values = append(values, cards.Seven)
	values = append(values, cards.Eight)
	values = append(values, cards.Nine)
	values = append(values, cards.Ten)
	values = append(values, cards.Jack)
	values = append(values, cards.Queen)
	values = append(values, cards.King)
	values = append(values, cards.Ace)

	for _, v := range values {
		card, err := cards.NewCard(v, suit)
		if err != nil {
			return nil, err
		}

		deck = append(deck, card)
	}

	return deck, nil
}
