package cardgame

import (
	"math/rand"
	"testing"
)

func TestPushAndPopToDeck(t *testing.T) {
	d := NewDeck()
	card := Card{Suit: 999}
	d.Push(card)
	act, err := d.Pop()
	if act != card {
		t.Errorf("not match act: %v != expected: %v", act, card)
	}
	if err != nil {
		t.Errorf("not expected err: %v", err)
	}
}

func TestReturnedErroAtPopWithEmptyDeck(t *testing.T) {
	d := NewDeck()
	act, err := d.Pop()
	if act != (Card{}) {
		t.Errorf("not match act: %v != expected: %v", act, Card{})
	}
	if err != ErrEmptyDeck {
		t.Errorf("not expected err: %v", err)
	}
}

func TestShuffle(t *testing.T) {
	d := NewDeck()
	d.Rand = rand.New(rand.NewSource(0))

	var startState, shuffledState []Card
	for i, n := range []int{11, 7, 6, 10, 5, 9, 2, 12, 4, 1, 8, 3, 13} {
		card := NormalCard(Spades, i+1)
		startState = append(startState, card)
		shuffledState = append(shuffledState, NormalCard(Spades, n))
		d.Push(card)
	}

	if !Equal(d.cards, startState) {
		t.Errorf("not match act: %v != startState: %v", d.cards, startState)
	}
	d.Shuffle()
	// d.ShuffleWithRand(rand.New(rand.NewSource(0)))
	if !Equal(d.cards, shuffledState) {
		t.Errorf("not match act: %v != shuffledState: %v", d.cards, shuffledState)
	}
}

func TestSort(t *testing.T) {
	d := NewDeck()
	d.Rand = rand.New(rand.NewSource(0))

	var startState, sortedState []Card
	for i, n := range []int{11, 7, 6, 10, 5, 9, 2, 12, 4, 1, 8, 3, 13} {
		card := NormalCard(Spades, n)
		startState = append(startState, card)
		sortedState = append(sortedState, NormalCard(Spades, i+1))
		d.Push(card)
	}

	if !Equal(d.cards, startState) {
		t.Errorf("not match act: %v != startState: %v", d.cards, startState)
	}
	d.Sort()
	// d.ShuffleWithRand(rand.New(rand.NewSource(0)))
	if !Equal(d.cards, sortedState) {
		t.Errorf("not match act: %v != shuffledState: %v", d.cards, sortedState)
	}
}
