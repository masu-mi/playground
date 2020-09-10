package blackjack

import (
	"fmt"
	"strings"

	"github.com/masu-mi/playground/training-code-design/cardgame-go"
)

type Choice int

const (
	limitScore = 21

	Hit Choice = 0 + iota
	Stand
)

type Game struct {
	deck    *cardgame.Deck
	players []*Player
}

func (c Choice) String() string {
	switch c {
	case Hit:
		return "Hit"
	case Stand:
		return "Stand"
	}
	panic(-1)
}

var _ cardgame.Game = &Game{}

func NewGame(ps ...*cardgame.Player) cardgame.Game {
	g := &Game{
		deck: cardgame.NewDeck(cardgame.NormalCards(6)),
	}
	registerPlayers(g, ps...)
	setupTable(g)
	return g
}

func registerPlayers(g *Game, ps ...*cardgame.Player) {
	for _, p := range ps {
		p := cliPlayer(p)
		g.players = append(g.players, p)
	}
}

func setupTable(b *Game) {
	b.deck.Shuffle()
	d := dealer()
	b.players = append(b.players, d)
	// ignore burst(because burst must not be caused)
	b.deal(2)
}

func (b *Game) Play() cardgame.Result {
	stopped := map[*Player]bool{}
	var standNum, burstedNum int
	lastOf := 0
	for true {
		b.displayTable()
		if r, stopped := b.judge(lastOf, standNum+burstedNum); stopped {
			b.displayOpenTable()
			return r
		}

		for _, p := range b.players {
			if stopped[p] {
				continue
			}
			c := p.Choice()
			if c == Stand {
				stopped[p] = true
				standNum++
				continue
			}
			b.dealTo(p)
			if p.hand.isBurst() {
				stopped[p] = true
				if p != b.dealer() {
					burstedNum++
				}
				continue
			}
		}
		lastOf++
	}
	// mock
	return cardgame.Result{}
}

func (b *Game) guests() []*Player {
	return b.players[0 : len(b.players)-1]
}

func (b *Game) judge(turn, stoppedNum int) (result cardgame.Result, finished bool) {
	r := cardgame.Result{Scores: map[int]int{}}
	dh := b.dealer().hand
	if turn == 0 && dh.isBlackjack() {
		for _, p := range b.guests() {
			if p.hand.isBlackjack() {
				r.Scores[p.id] = 0
			} else {
				r.Scores[p.id] = -1
			}
		}
		return r, true
	}
	if stoppedNum < len(b.players)-1 {
		return cardgame.Result{}, false
	}
	if dh.isBurst() {
		for _, p := range b.guests() {
			if p.hand.isBurst() {
				r.Scores[p.id] = 0
			} else {
				r.Scores[p.id] = 1
			}
		}
		return r, true
	}
	dealerScore := dh.value()
	for _, p := range b.guests() {
		if p.hand.isBurst() {
			r.Scores[p.id] = -1
		} else if pv := p.hand.value(); pv < dealerScore {
			r.Scores[p.id] = -1
		} else if pv == dealerScore {
			r.Scores[p.id] = 0
		} else {
			r.Scores[p.id] = 1
		}
	}
	return r, true
}

type pair struct {
	f, s int
	h    hand
}
type ranking []pair

func (r ranking) Len() int {
	return len(r)
}

func (r ranking) Less(i, j int) bool {
	return r[i].s < r[j].s
}

func (r ranking) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (b *Game) displayTable() {
	fmt.Println("+-------------------------------+")
	for _, p := range b.players[0 : len(b.players)-1] {
		fmt.Printf(" %s: %v\n", p.name, p.hand)
	}
	fmt.Printf(" dealer: [**], %v\n", b.dealer().hand[1:])
	fmt.Println("+-------------------------------+")
}
func (b *Game) displayOpenTable() {
	fmt.Println("+-------------------------------+")
	for _, p := range b.players {
		fmt.Printf(" %s(%d): %v\n", p.name, p.hand.value(), p.hand)
	}
	fmt.Println("+-------------------------------+")
}

func (b *Game) dealer() *Player {
	return b.players[len(b.players)-1]
}

func (b *Game) deal(n int) (num int) {
	for i := 0; i < n; i++ {
		for _, p := range b.players {
			if b.dealTo(p) {
				num++
			}
		}
	}
	return num
}

func (b *Game) dealTo(p *Player) (bursted bool) {
	h := p.hand
	// ignore error handle
	c, _ := b.deck.Pop()
	h = append(h, c)
	p.hand = h
	return h.isBurst()
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

func (h hand) String() string {
	b := &strings.Builder{}
	b.WriteString(fmt.Sprintf("%v", h[0]))
	for i := 1; i < len(h); i++ {
		b.WriteString(", ")
		b.WriteString(fmt.Sprintf("%v", h[i]))
	}
	return b.String()
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
