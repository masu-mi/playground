package main

import "testing"

func Test_sortAsc(t *testing.T) {
	act := []byte("cba")
	sortAsc(act)
	if !equalList(act, []byte("abc")) {
		t.Errorf("don't match, expected: %s; act: %s", "abc", act)
	}
}
func Test_sortDesc(t *testing.T) {
	act := []byte("abc")
	sortDesc(act)
	if !equalList(act, []byte("cba")) {
		t.Errorf("don't match, expected: %s; act: %s", "abc", act)
	}
}

func equalList(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
