package main

import (
	"testing"
)

func TestDigit(t *testing.T) {
	type testCase struct{ i, e int }
	for _, c := range []testCase{
		testCase{0, 0},
		testCase{1, 1},
		testCase{9, 1},
		testCase{10, 2},
	} {
		got := digit(c.i)
		want := c.e
		if got != want {
			t.Fatalf("want %v, but %v:", want, got)
		}
	}
}
func TestDigitWithBase(t *testing.T) {
	type testCase struct{ i, b, e int }
	for _, c := range []testCase{
		testCase{0, 2, 0},
		testCase{1, 2, 1},
		testCase{2, 2, 2},
		testCase{8, 2, 4},
		testCase{9, 2, 4},
		testCase{10, 2, 4},
		testCase{8, 8, 2},
	} {
		got := digitWithBase(c.i, c.b)
		want := c.e
		if got != want {
			t.Fatalf("want %v, but %v:", want, got)
		}
	}
}
