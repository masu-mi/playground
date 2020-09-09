package blackjack

import (
	"fmt"

	"github.com/masu-mi/playground/training-code-design/cardgame-go"
)

type Player interface {
	ID() int
	Name() string
	Choice(c []cardgame.Card) Choice
}

var _ Player = &CliPlayer{}

type CliPlayer struct {
	*cardgame.Player
}

func (c *CliPlayer) Choice(cards []cardgame.Card) Choice {
	// fmt.Printf("User(%s): %v\n", c.Player.Name, cards)
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
func (cp *CliPlayer) Name() string {
	return cp.Player.Name
}
func (cp *CliPlayer) ID() int {
	return cp.Player.ID
}

type dealer struct{}

var _ Player = &dealer{}

func (c *dealer) Choice(cards []cardgame.Card) Choice {
	if hand(cards).willDealerStand() {
		return Stand
	}
	return Hit
}
func (cp *dealer) Name() string {
	return "dealer"
}
func (cp *dealer) ID() int {
	return -1
}
