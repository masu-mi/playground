package main

import (
	"testing"
)

func TestDivisorNumIs(t *testing.T) {
	got := divisorNumIs(25, 3)
	want := true
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
