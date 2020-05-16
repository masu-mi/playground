package unionfind

import (
	"testing"
)

func TestInit(t *testing.T) {
	size := 5
	uf := newUnifonFind(size)
	for i := 0; i < size; i++ {
		if got := uf.find(i); got != i {
			t.Fatalf("want %d, but %d", i, got)
		}
	}
}

func TestSame(t *testing.T) {
	size := 5
	uf := newUnifonFind(size)
	uf.union(0, 1)
	uf.union(1, 2)
	uf.union(3, 4)

	type testCase struct {
		x, y   int
		result bool
	}
	for _, c := range []testCase{
		testCase{0, 0, true},
		testCase{1, 1, true},
		testCase{2, 2, true},
		testCase{3, 3, true},
		testCase{4, 4, true},
		testCase{0, 1, true},
		testCase{1, 2, true},
		testCase{0, 2, true},
		testCase{3, 4, true},
		testCase{0, 3, false},
	} {
		got := uf.same(c.x, c.y)
		if got != c.result {
			t.Fatalf("want %t, but %t", c.result, got)
		}
	}
}

func TestConnected(t *testing.T) {
	size := 3

	type unify struct{ x, y int }
	type testCase struct {
		unify  []unify
		result bool
	}
	for _, c := range []testCase{
		testCase{
			[]unify{
				unify{0, 1},
				unify{1, 2},
			},
			true,
		},
		testCase{
			[]unify{
				unify{0, 1},
			},
			false,
		},
	} {
		uf := newUnifonFind(size)
		for _, unify := range c.unify {
			uf.union(unify.x, unify.y)
		}
		got := uf.connected
		if got != c.result {
			t.Fatalf("want %t, but %t", c.result, got)
		}
	}
}
