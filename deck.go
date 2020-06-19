//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit of a card
type Suit uint8

// Different card suits
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank of a card
type Rank uint8

// Different card ranks
const (
	Joker Rank = iota
	Ace
	Two
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
)

// Min and max ranks
const (
	MinRank = Ace
	MaxRank = King
)

// Card with rank and suit
type Card struct {
	Rank    Rank
	Suit    Suit
	Visible bool
	Point   int
}

// Points returns sum of points from provided cards
func Points(cards *[]Card) int {
	res := 0
	for _, card := range *cards {
		if card.Visible {
			res += card.Point
		}
	}
	return res
}

func (c Card) String() string {
	if !c.Visible {
		return "FACEDOWN"
	}
	if c.Rank == Joker {
		return Joker.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// Deck is constructed from multiple cards
type Deck []Card

// Draw draws first n cards from the deck and removes those cards for it
func (d *Deck) Draw(n int) []Card {
	cards := []Card{}
	for i := 0; i < n; i++ {
		cards = append(cards, (*d)[i])
	}
	*d = (*d)[n:]
	return cards
}

// FacePoints sets n points to all Face cards (Jack, Queen, King)
func (d *Deck) FacePoints(n int) {
	deck := *d
	for i, card := range deck {
		if card.Rank == Jack || card.Rank == Queen || card.Rank == King {
			deck[i].Point = n
		}
	}
}

// RankPoints sets n points for all cards of rank r
func (d *Deck) RankPoints(r Rank, n int) {
	deck := *d
	for i, card := range deck {
		if card.Rank == r {
			deck[i].Point = n
		}
	}
}

// SuitPoints sets n points for all cards of suit s
func (d *Deck) SuitPoints(s Suit, n int) {
	deck := *d
	for i, card := range deck {
		if card.Suit == s {
			deck[i].Point = n
		}
	}
}

// AddPoints sets n points to a card with rank r and suir s
func (d *Deck) AddPoints(r Rank, s Suit, n int) {
	deck := *d
	for i, card := range deck {
		if r == Joker && card.Rank == r {
			deck[i].Point = n
		} else if card.Rank == r && card.Suit == s {
			deck[i].Point = n
		}
	}
}

// Option represents a type for functional options
type Option func(d *Deck)

// New creates a new deck of playing cards
func New(opts ...Option) Deck {
	deck := Deck{}
	for _, suit := range suits {
		for i := MinRank; i <= MaxRank; i++ {
			deck = append(deck, Card{Rank: i, Suit: suit, Visible: true})
		}
	}
	defaultPoints(&deck)
	for _, opt := range opts {
		opt(&deck)
	}
	return deck
}

// Sort with custom Less fucntion
func Sort(less func(d *Deck) func(i, j int) bool) Option {
	return func(d *Deck) {
		sort.Slice(*d, less(d))
	}
}

// DefaultSort sorts by suit in ascending order
func DefaultSort() Option {
	return func(d *Deck) {
		sort.Slice(*d, less(d))
	}
}

func less(d *Deck) func(i, j int) bool {
	return func(i, j int) bool {
		return defSort((*d)[i]) < defSort((*d)[j])
	}
}

func defSort(c Card) int {
	return int(c.Suit)*int(MaxRank) + int(c.Rank)
}

// Shuffle shuffles the deck
func Shuffle() Option {
	return func(d *Deck) {
		deck := *d
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	}
}

// Jokers adds n number of jokers to the deck
func Jokers(n int) Option {
	return func(d *Deck) {
		for i := 0; i < n; i++ {
			*d = append(*d, Card{Rank: Joker, Suit: 0})
		}
	}
}

// Without creates a deck without specified cards
func Without(ranks ...Rank) Option {
	return func(d *Deck) {
		deck := Deck{}
		for _, v := range *d {
			if contains(&ranks, v.Rank) {
				continue
			}
			deck = append(deck, Card{Rank: v.Rank, Suit: v.Suit})
		}
		*d = deck
	}
}

func contains(ranks *[]Rank, rank Rank) bool {
	for _, v := range *ranks {
		if v == rank {
			return true
		}
	}
	return false
}

// Size n determens of how many decks will the deck be formed
func Size(n int) Option {
	return func(d *Deck) {
		for i := 1; i < n; i++ {
			*d = append(*d, New()...)
		}
	}
}

func defaultPoints(d *Deck) {
	deck := *d
	for i, card := range deck {
		if card.Rank == Jack || card.Rank == Queen || card.Rank == King {
			deck[i].Point = int(card.Rank) + 1
		} else {
			deck[i].Point = int(card.Rank)
		}
	}
}
