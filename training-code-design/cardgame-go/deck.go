package cardgame

import (
	"errors"
	"math/rand"
	"sort"
	"time"
)

// Deck : Should the structure support to PushBottom, GetCardAt, Cut, GetSnapshot?
type Deck struct {
	*rand.Rand
	number int
	cards  []Card
}

// NewDeck returns empty Deck
func NewDeck(cards []Card) *Deck {
	t := time.Now()
	return &Deck{
		Rand:   rand.New(rand.NewSource(t.UnixNano())),
		cards:  cards,
		number: len(cards),
	}
}

// ErrEmptyDeck is a error
var ErrEmptyDeck = errors.New("empty deck")

func (d *Deck) Push(c Card) {
	d.cards = append(d.cards, c)
	d.number++
}

func (d *Deck) Pop() (c Card, e error) {
	if d.number <= 0 {
		return Card{}, ErrEmptyDeck
	}
	d.number--
	c, d.cards = d.cards[d.number], d.cards[0:d.number]
	return c, nil
}

func (d *Deck) Shuffle() {
	// Is receiver metho correct?
	d.ShuffleWithRand(d.Rand)
}

func (d *Deck) ShuffleWithRand(r *rand.Rand) {
	r.Shuffle(d.number, func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Sort() {
	// Is receiver metho correct?
	d.SortWithComp(SimpleLess)
}

func (d *Deck) SortWithComp(comp func(i, j Card) bool) {
	s := sortableCards{
		cards: d.cards,
		less:  comp,
	}
	sort.Sort(s)
}

type sortableCards struct {
	cards []Card
	less  func(i, j Card) bool
}

func (s sortableCards) Len() int {
	return len(s.cards)
}

func (s sortableCards) Less(i, j int) bool {
	return s.less(s.cards[i], s.cards[j])
}

func (s sortableCards) Swap(i, j int) {
	s.cards[i], s.cards[j] = s.cards[j], s.cards[i]
}
