package deck

import (
	"fmt"
	"testing"
)

var standardDeckCount int = 4 * 13

func ExampleCard() {
	fmt.Println(Card{Rank: King, Suit: Diamond, Visible: true})
	fmt.Println(Card{Rank: Ace, Suit: Spade, Visible: true})
	fmt.Println(Card{Rank: Ten, Suit: Heart, Visible: true})
	fmt.Println(Card{Rank: Jack, Suit: Club, Visible: true})
	fmt.Println(Card{Rank: Joker, Suit: Diamond, Visible: true})
	fmt.Println(Card{Rank: Jack, Suit: Club, Visible: false})
	fmt.Println(Card{Rank: Jack, Suit: Club})

	// Output:
	// | King of Diamonds |
	// | Ace of Spades |
	// | Ten of Hearts |
	// | Jack of Clubs |
	// | Joker |
	// | FACE DOWN |
	// | FACE DOWN |
}

func TestNew(t *testing.T) {
	deck := New()
	if len(deck) != standardDeckCount {
		t.Errorf("Error running a New() method, number of cards created is %d expected %d", len(deck), standardDeckCount)
	}
}

func TestSort(t *testing.T) {
	sortedDeck := New(Sort(descending))
	deck := New()
	if sortedDeck[0] != deck[len(deck)-1] {
		t.Errorf("Error in TestSort, %s isn't equal to %s", sortedDeck[0], deck[len(deck)-1])
	}
}

func descending(d *Deck) func(i, j int) bool {
	return func(i, j int) bool {
		return defSort((*d)[i]) > defSort((*d)[j])
	}
}

func TestDefaultSort(t *testing.T) {
	sortedDeck := New(DefaultSort())
	card := Card{Rank: Ace, Suit: Spade, Visible: true}
	if sortedDeck[0] != card {
		t.Errorf("First card was expected to be %s but it was %s", card, sortedDeck[0])
	}
}

func TestShuffle(t *testing.T) {
	deck := New()
	shuffleDeck := New(Shuffle())
	if len(deck) != len(shuffleDeck) {
		t.Errorf("Error while Shuffling, length of shuffled deck is %d while it should be %d", len(shuffleDeck), len(deck))
	}
	size := len(deck)
	same := 0
	for i := 0; i < size; i++ {
		if deck[i] == shuffleDeck[i] {
			same++
		}
	}
	if same == size {
		t.Error("Shuffled deck has the same card order as standard deck. Shuffling didn't work")
	}
}

func TestJokers(t *testing.T) {
	n := 3
	deck := New(Jokers(n))
	count := 0
	for _, card := range deck {
		if card.Rank == Joker {
			count++
		}
	}
	if count != n {
		t.Errorf("Error while adding Jokers to the deck, expercted %d, added %d", n, count)
	}
}

func TestWithout(t *testing.T) {
	deck := New(Without(Ace, Eight, Ten))
	for _, card := range deck {
		if card.Rank == Ace || card.Rank == Eight || card.Rank == Ten {
			t.Errorf("Error removing ranks from deck. %s is expected to be removed, but it exists in the deck.", card.Rank)
		}
	}
}

func TestSize(t *testing.T) {
	n := 5
	deck := New(Size(n))
	if standardDeckCount*n != len(deck) {
		t.Errorf("Error setting deck size. Size expected to be %d but it was %d", standardDeckCount*n, len(deck))
	}

}
