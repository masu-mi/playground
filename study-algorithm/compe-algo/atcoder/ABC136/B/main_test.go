package main

import (
	"testing"
)

func TestDigit(t *testing.T) {
	type testCase struct {
		input, expected int
	}
	for _, c := range []testCase{
		testCase{0, 0},
		testCase{1, 1},
		testCase{9, 1},
		testCase{10, 2},
		testCase{99, 2},
		testCase{100, 3},
	} {
		got := digit(c.input)
		if got != c.expected {
			t.Fatalf("want %v, but %v:", c.expected, got)
		}
	}
}
