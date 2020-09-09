package blackjack

import (
	"fmt"
	"sort"
	"strings"

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
	deck    *cardgame.Deck
	players []Player
	hands   map[Player]hand
}

var _ cardgame.Game = &Blackjack{}

func NewBlackjack() *Blackjack {
	return &Blackjack{
		deck:  cardgame.NewDeck(cardgame.NormalCards(6)),
		hands: map[Player]hand{},
	}
}

func (b *Blackjack) RegisterPlayers(ps ...*cardgame.Player) {
	for _, p := range ps {
		p := &CliPlayer{Player: p}
		b.players = append(b.players, p)
		b.hands[p] = hand{}
	}
}

func (b *Blackjack) setupTable() {
	b.deck.Shuffle()
	d := &dealer{}
	b.players = append(b.players, d)
	b.hands[d] = hand{}
	// ignore burst(because burst must not be caused)
	b.deal(2)
}

func (b *Blackjack) Play() cardgame.Result {
	stopped := map[Player]bool{}
	b.setupTable()
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
			c := p.Choice(b.hands[p])
			if c == Stand {
				stopped[p] = true
				standNum++
				continue
			}
			b.dealTo(p)
			if b.hands[p].isBurst() {
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

func (b *Blackjack) playerRanking() []pair {
	var ps []pair
	for _, player := range b.players[:len(b.players)-1] {
		h := b.hands[player]
		ps = append(ps, pair{f: player.ID(), s: h.value(), h: h})
	}
	sort.Sort(sort.Reverse(ranking(ps)))
	return ps
}

func (b *Blackjack) judge(turn, stoppedNum int) (result cardgame.Result, finished bool) {
	ranking := b.playerRanking()
	r := cardgame.Result{Scores: map[int]int{}}
	dh := b.dealerHand()
	if turn == 0 && dh.isBlackjack() {
		for _, p := range ranking {
			if p.h.isBlackjack() {
				r.Scores[p.f] = 0
			} else {
				r.Scores[p.f] = -1
			}
		}
		return r, true
	}
	if stoppedNum < len(b.players)-1 {
		return cardgame.Result{}, false
	}
	if dh.isBurst() {
		for _, p := range ranking {
			if p.h.isBurst() {
				r.Scores[p.f] = 0
			} else {
				r.Scores[p.f] = 1
			}
		}
		return r, true
	}
	dealerScore := dh.value()
	for _, p := range ranking {
		if p.h.isBurst() {
			r.Scores[p.f] = -1
		} else if pv := p.h.value(); pv < dealerScore {
			r.Scores[p.f] = -1
		} else if pv == dealerScore {
			r.Scores[p.f] = 0
		} else {
			r.Scores[p.f] = 1
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

func (b *Blackjack) displayTable() {
	fmt.Println("+-------------------------------+")
	for _, p := range b.players[0 : len(b.players)-1] {
		fmt.Printf(" %s: %v\n", p.Name(), b.hands[p])
	}
	fmt.Printf(" dealer: [**], %v\n", b.dealerHand()[1:])
	fmt.Println("+-------------------------------+")
}
func (b *Blackjack) displayOpenTable() {
	fmt.Println("+-------------------------------+")
	for _, p := range b.players {
		fmt.Printf(" %s(%d): %v\n", p.Name(), b.hands[p].value(), b.hands[p])
	}
	fmt.Println("+-------------------------------+")
}

func (b *Blackjack) dealer() Player {
	return b.players[len(b.players)-1]
}

func (b *Blackjack) dealerHand() hand {
	return b.hands[b.dealer()]
}

func (b *Blackjack) deal(n int) (num int) {
	for i := 0; i < n; i++ {
		for _, p := range b.players {
			if b.dealTo(p) {
				num++
			}
		}
	}
	return num
}

func (b *Blackjack) dealTo(p Player) (bursted bool) {
	h := b.hands[p]
	// ignore error handle
	c, _ := b.deck.Pop()
	h = append(h, c)
	b.hands[p] = h
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
