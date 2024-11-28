package maumau

type Card struct {
	Suit string
	Rank string
}

func NewCard(suit string, rank string) Card {
	return Card{Suit: suit, Rank: rank}
}
