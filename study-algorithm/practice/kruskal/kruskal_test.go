package kruskal

import (
	"testing"
)

func TestFindMSTWithKruskal(t *testing.T) {
	type testCase struct {
		card  int
		edges []edge
		cost  int
	}
	for _, c := range []testCase{
		testCase{
			card:  1,
			edges: []edge{},
			cost:  0,
		},
		testCase{
			card:  2,
			edges: []edge{edge{0, 1, 0}},
			cost:  0,
		},
		testCase{
			card:  2,
			edges: []edge{edge{0, 1, 1}},
			cost:  1,
		},
		testCase{
			card: 3,
			edges: []edge{
				edge{0, 1, 1},
				edge{0, 2, 5},
				edge{2, 1, 2},
			},
			cost: 3,
		},
		testCase{
			card: 4,
			edges: []edge{
				edge{0, 1, 1},
				edge{0, 2, 5},
				edge{2, 1, 2},
				edge{2, 3, 10},
				edge{0, 3, 2},
			},
			cost: 5,
		},
		testCase{
			card: 4,
			edges: []edge{
				edge{0, 1, 4},
				edge{0, 2, 1},
				edge{1, 2, 4},
				edge{1, 3, 2},
				edge{2, 3, 1},
			},
			cost: 4,
		},
	} {
		_, cost, err := findMSTWithKruskal(c.card, c.edges)
		if err != nil {
			t.Fatalf("error retunred %v", err)
		}
		if cost != c.cost {
			t.Fatalf("want %d, but %d", c.cost, cost)
		}
	}
}
