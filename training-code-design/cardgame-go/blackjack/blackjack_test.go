package blackjack

import (
	"testing"

	"github.com/masu-mi/playground/training-code-design/cardgame-go"
)

func TestHandValue(t *testing.T) {
	type testCase struct {
		input    []int
		expected int
	}

	for idx, c := range []testCase{
		testCase{input: []int{1, 10}, expected: 21},
		testCase{input: []int{2, 2}, expected: 4},
		testCase{input: []int{10, 10, 10}, expected: 30},
		testCase{input: []int{1, 1, 10}, expected: 12},
	} {
		var inputs []cardgame.Card
		for _, num := range c.input {
			inputs = append(inputs, cardgame.NormalCard(cardgame.Spades, num))
		}
		if act := hand(inputs).value(); act != c.expected {
			t.Errorf("idx:%d, expected [%v] but got %v", idx, c.expected, act)
		}
	}
}
