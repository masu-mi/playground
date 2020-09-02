package cardgame

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

func NewNormalCard(s int, n int) Card {
	return Card{Suit: s, Number: n}
}
func NewSpecialCard(name string) Card {
	return Card{SpecialName: name}
}
func NewJoker() Card {
	return NewSpecialCard(joker)
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
