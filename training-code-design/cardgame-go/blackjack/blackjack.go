package blackjack

import (
	"fmt"

	"github.com/masu-mi/playground/training-code-design/cardgame-go"
)

type Choice int

const (
	limitScore = 21

	Hit Choice = 0 + iota
	Stand
)

func (c Choice) String() string {
	switch c {
	case Hit:
		return "Hit"
	case Stand:
		return "Stand"
	}
	panic(-1)
}

type Blackjack struct {
	dealer  hand
	players []Player
	hands   map[Player]hand
}

var _ cardgame.Game = &Blackjack{}

func NewBlackjack() *Blackjack {
	return &Blackjack{hands: map[Player]hand{}}
}

type Player interface {
	Choice(c []cardgame.Card) Choice
}

type CliPlayer struct {
	*cardgame.Player
}

func (c *CliPlayer) Choice(cards []cardgame.Card) Choice {
	fmt.Printf("User(%s): %v\n", c.Player.Name, cards)
	fmt.Println("Hit/Stand [H/S]?")
	var choice string
	fmt.Scanf("%s", &choice)
	switch choice {
	case "H", "h":
		return Hit
	case "S", "s":
		return Stand
	}
	panic(-1)
}

func (b *Blackjack) RegisterPlayers(ps ...*cardgame.Player) {
	for _, p := range ps {
		p := &CliPlayer{Player: p}
		b.players = append(b.players, p)
		b.hands[p] = hand{}
	}
}

func (b *Blackjack) Play() cardgame.Result {
	for _, p := range b.players {
		c := p.Choice(b.hands[p])
		fmt.Println(c)
	}
	// mock
	return cardgame.Result{}
}

type hand []cardgame.Card

func (h hand) value() int {
	aces, maxScore := 0, 0
	for i := 0; i < len(h); i++ {
		switch num := h[i].Number; num {
		case 1: // ace can select 1 or 11 points
			aces++
			maxScore += 11
		case 11, 12, 13: // picture card is 10 points
			maxScore += 10
		default:
			maxScore += num
		}
	}
	// like ceil func
	select1Num := min(aces, (max(maxScore-limitScore, 0)+9)/10)
	return maxScore - 10*select1Num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (h hand) isBurst() bool {
	return h.value() > limitScore
}

func (h hand) isBlackjack() bool {
	return h.value() == limitScore
}

func (h hand) willDealerStand() bool {
	return h.value() >= 17
}
