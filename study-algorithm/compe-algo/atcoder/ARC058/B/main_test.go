package main

import (
	"testing"
)

func TestModuloLog(t *testing.T) {
	type testCase struct{ a, b, modulo, exp int }
	for _, c := range []testCase{
		testCase{a: 2, b: 2, modulo: 3, exp: 1},
		testCase{a: 2, b: 1, modulo: 3, exp: 0},
		testCase{a: 2, b: 1, modulo: 5, exp: 0},
		testCase{a: 2, b: 2, modulo: 5, exp: 1},
		testCase{a: 2, b: 3, modulo: 5, exp: 3},
		testCase{a: 2, b: 4, modulo: 5, exp: 2},
		testCase{a: 2, b: 0, modulo: 5, exp: -1},
		testCase{a: 2, b: 3, modulo: 11, exp: 8},
	} {
		got := moduloLog(c.a, c.b, c.modulo)
		if got != c.exp {
			t.Fatalf("want %v, but %v:", c.exp, got)
		}
	}
}
