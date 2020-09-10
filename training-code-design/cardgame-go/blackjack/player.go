package blackjack

import (
	"fmt"

	"github.com/masu-mi/playground/training-code-design/cardgame-go"
)

type cmdChoice func(c []cardgame.Card) Choice

type Player struct {
	id     int
	name   string
	hand   hand
	choice cmdChoice
}

func (p *Player) Choice() Choice {
	return p.choice(p.hand)
}

func cliPlayer(p *cardgame.Player) *Player {
	return &Player{
		id:     p.ID,
		name:   p.Name,
		choice: choiceFromStdin,
	}
}

func choiceFromStdin(cards []cardgame.Card) Choice {
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

func dealer() *Player {
	return &Player{
		id:     -1,
		name:   "dealer",
		choice: choiceAsDealer,
	}
}

func choiceAsDealer(cards []cardgame.Card) Choice {
	if hand(cards).willDealerStand() {
		return Stand
	}
	return Hit
}
