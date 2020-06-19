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
	// King of Diamonds
	// Ace of Spades
	// Ten of Hearts
	// Jack of Clubs
	// Joker
	// FACEDOWN
	// FACEDOWN
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
	card := Card{Rank: Ace, Suit: Spade, Visible: true, Point: 1}
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

func TestDraw(t *testing.T) {
	nToDraw := 3
	deck := New()
	deckLength := len(deck)
	cards := deck.Draw(nToDraw)
	firstNFromDeck := New()[:nToDraw]
	for i, card := range cards {
		if card != firstNFromDeck[i] {
			t.Errorf("Error drawing the card on possition %d. Expected %s, got %s", i, firstNFromDeck[i], card)
		}
	}
	if deckLength-nToDraw != len(deck) {
		t.Errorf("Error while removing cards from the deck after drawing. Expected length of a new deck %d, got %d", deckLength-nToDraw, len(deck))
	}
}

func TestDefaultPoints(t *testing.T) {
	deck := New()
	for _, card := range deck {
		if card.Rank == Jack || card.Rank == Queen || card.Rank == King {
			if card.Point != int(card.Rank)+1 {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, int(card.Rank)+1)
			}
		} else {
			if card.Point != int(card.Rank) {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, int(card.Rank))
			}
		}
	}
}

func TestFacePoints(t *testing.T) {
	facePoints := 10
	deck := New()
	deck.FacePoints(facePoints)
	for _, card := range deck {
		if card.Rank == Jack || card.Rank == Queen || card.Rank == King {
			if card.Point != facePoints {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, facePoints)
			}
		} else {
			if card.Point != int(card.Rank) {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, int(card.Rank))
			}
		}
	}
}

func TestRankPoints(t *testing.T) {
	rankPts := 10
	deck := New()
	chosenRank := Five
	deck.RankPoints(chosenRank, rankPts)
	for _, card := range deck {
		if card.Rank == chosenRank {
			if card.Point != rankPts {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, rankPts)
			}
		} else if card.Rank == Jack || card.Rank == Queen || card.Rank == King {
			if card.Point != int(card.Rank)+1 {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, int(card.Rank)+1)
			}
		} else {
			if card.Point != int(card.Rank) {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, int(card.Rank))
			}
		}
	}
}

func TestSuitPoints(t *testing.T) {
	suitPts := 10
	deck := New()
	chosenSuit := Heart
	deck.SuitPoints(chosenSuit, suitPts)
	for _, card := range deck {
		if card.Suit == chosenSuit {
			if card.Point != suitPts {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, suitPts)
			}
		}
	}
}

func TestAddPoints(t *testing.T) {
	pts := 75
	deck := New()
	chosenSuit := Heart
	chosenRank := Queen
	deck.AddPoints(chosenRank, chosenSuit, pts)
	for _, card := range deck {
		if card.Suit == chosenSuit && card.Rank == chosenRank {
			if card.Point != pts {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, pts)
			}
		}
	}
	chosenRank = Joker
	jokerDeck := New(Jokers(5))
	jokerPts := 90
	jokerDeck.AddPoints(Joker, Diamond, jokerPts)
	for _, card := range deck {
		if card.Rank == Joker {
			if card.Point != jokerPts {
				t.Errorf("Rank for %s is %d, expected %d", card, card.Point, jokerPts)
			}
		}
	}
}

func TestPoints(t *testing.T) {
	d := New()
	cards := d.Draw(13)
	pts := Points(&cards)
	if pts != 94 {
		t.Error("Wrong sum of cards")
	}
}
