package cardgame

import "fmt"

const (
	forSpecialCards = 0 + iota

	Clubs
	Diamonds
	Hearts
	Spades
)

const (
	joker = "Joker"
	// ExtraJoker = "ExtraJoker"
)

type Card struct {
	Suit, Number int
	SpecialName  string
}

func NormalCard(s int, n int) Card {
	return Card{Suit: s, Number: n}
}
func SpecialCard(name string) Card {
	return Card{SpecialName: name}
}
func Joker() Card {
	return SpecialCard(joker)
}

func NormalCards(set int) (cards []Card) {
	for i := 0; i < set; i++ {
		for j := 1; j <= 13; j++ {
			for _, s := range []int{Clubs, Diamonds, Hearts, Spades} {
				cards = append(cards, NormalCard(s, j))
			}
		}
	}
	return cards
}

func (c Card) String() string {
	if c.IsSpecial() {
		return fmt.Sprintf("[%s]", c.SpecialName)
	}
	var s string
	switch c.Suit {
	case Clubs:
		s = "♣︎"
	case Diamonds:
		s = "♢"
	case Hearts:
		s = "♡"
	case Spades:
		s = "♠︎"
	}
	var num string
	switch c.Number {
	case 11:
		num = "J"
	case 12:
		num = "Q"
	case 13:
		num = "K"
	case 1:
		num = "A"
	default:
		num = fmt.Sprintf("%d", c.Number)
	}
	return fmt.Sprintf("[%s,%s]", s, num)
}

func (c Card) Valid() bool {
	return c.IsNormal() || c.IsSpecial()
}

func (c Card) IsNormal() bool {
	if c.SpecialName != "" {
		return false
	}
	if c.Suit < Clubs || c.Suit > Spades {
		return false
	}
	if c.Number < 1 || c.Number > 13 {
		return false
	}
	return true
}

func (c Card) IsSpecial() bool {
	return c.Suit == 0 && c.SpecialName != ""
}

func Equal(a, b []Card) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func SimpleLess(i, j Card) bool {
	if i.Number < j.Number {
		return true
	}
	if i.Number == j.Number {
		return i.Suit < j.Suit
	}
	return false
}
