package main

import (
	"fmt"
	"math"
	"testing"
)

func TestLST_Max(t *testing.T) {
	type q struct{ a, b, v int }
	type testCase struct {
		name  string
		last  int
		setup []q
		get   []q
	}
	for _, c := range []testCase{
		testCase{
			name: "normal 1 item",
			last: 1,
			setup: []q{
				q{0, 1, 1},
			},
			get: []q{
				q{0, 0, nINF},
				q{1, 1, nINF},
				q{0, 1, 1},
				q{1, 2, 0},
			},
		},
		testCase{
			name: "multipl update 1 item",
			last: 1,
			setup: []q{
				q{0, 1, 1},
				q{0, 1, 2},
				q{0, 1, 4},
				q{0, 1, 5},
				q{0, 1, -2},
			},
			get: []q{
				q{0, 0, nINF},
				q{1, 1, nINF},
				q{0, 1, 10},
			},
		},
		testCase{
			name: "range update",
			last: 10,
			setup: []q{
				q{0, 5, 1},
				q{2, 6, 2},
				q{4, 8, 4},
				q{6, 10, -1},
				q{10, 11, 100},
			},
			get: []q{
				q{0, 1, 1},
				q{1, 2, 1},
				q{2, 3, 3},
				q{3, 4, 3},
				q{4, 5, 7},
				q{5, 6, 6},
				q{6, 7, 3},
				q{7, 8, 3},
				q{8, 9, -1},
				q{9, 10, -1},
				q{10, 11, 100},
			},
		},
		testCase{
			name: "normal 2 item",
			last: 5,
			setup: []q{
				q{0, 1, 1},
				q{3, 4, 2},
				q{5, 6, 1},
			},
			get: []q{
				q{0, 1, 1},
				q{0, 2, 1},
				q{0, 3, 1},
				q{0, 4, 2},
				q{0, 5, 2},
				q{0, 6, 2},
				q{1, 6, 2},
				q{2, 6, 2},
				q{3, 6, 2},
				q{4, 6, 1},
				q{5, 6, 1},
				q{6, 6, nINF},
				q{6, 7, nINF},
			},
		},
		testCase{
			name: "negative 1 item",
			last: 1,
			setup: []q{
				q{0, 1, -1},
			},
			get: []q{
				q{0, 0, nINF},
				q{1, 1, nINF},
				q{0, 1, -1},
				q{1, 2, 0},
				q{2, 2, nINF},
				q{2, 3, nINF},
			},
		},
		testCase{
			name: "invalid out of target",
			last: 1,
			setup: []q{
				q{0, 1, 1},
			},
			get: []q{
				q{-1, 0, nINF},
				q{2, 2, nINF},
				q{2, 3, nINF},
			},
		},
		testCase{
			name: "multi write",
			last: 8,
			setup: []q{
				q{0, 1, 1},
				q{1, 3, 1},
				q{2, 4, 5},
				q{3, 4, -2},
				q{5, 8, 10},
				q{6, 7, -4},
			},
			get: []q{
				q{0, 2, 1},
				q{1, 3, 6},
				q{4, 5, 0},
				q{6, 7, 6},
				q{7, 9, 10},
			},
		},
	} {
		tr := newLazySegTree(c.last)
		tr.update(0, c.last+1, -tr.max(0, 1))
		for _, q := range c.setup {
			tr.update(q.a, q.b, q.v)
		}
		for idx, q := range c.get {
			t.Run(fmt.Sprintf("%s (%d)", c.name, idx), func(t *testing.T) {
				got := tr.max(q.a, q.b)
				if got != q.v {
					fmt.Printf("%v\n", tr.mon[tr.size-1:])
					t.Fatalf("want %v, but %v:", q.v, got)
				}
			})
		}
	}
}

var (
	nINF = math.MinInt32
)
