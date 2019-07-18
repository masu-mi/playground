package main

import (
	"fmt"
)

type state [3][]int

func main() {

	var n, m int
	fmt.Scan(&n, &m)

	start := state{}
	for i := 0; i < 3; i++ {
		var c int
		fmt.Scan(&c)
		for j := 0; j < c; j++ {
			var num int
			fmt.Scan(&num)
			start[i] = append(start[i], num)
		}
	}
	fmt.Printf("%d\n", opNum(start, n, m))
}

type procState struct {
	s state
	n int
}

type op struct {
	from, to int
}

func opNum(start state, n, limit int) int {
	var q []procState
	q = append(q, procState{start, 0})
	for len(q) > 0 {
		cur := q[0]
		q = q[1:len(q)]
		if reached(cur.s, n) {
			return cur.n
		}
		if cur.n > limit {
			continue
		}
		for _, o := range []op{{0, 1}, {1, 0}, {1, 2}, {2, 1}} {
			// A->B, B->A, B->C, C->B
			if !movable(cur.s, o.from, o.to) {
				continue
			}
			next := movedState(cur.s, o.from, o.to)
			q = append(q, procState{next, cur.n + 1})
		}
	}
	return -1
}

func movedState(s state, from, to int) state {
	next := state{}
	for i, stack := range s {
		next[i] = append(next[i], stack...)
	}
	var top int
	top, next[from] = next[from][len(next[from])-1], next[from][0:len(next[from])-1]
	next[to] = append(next[to], top)
	return next
}

func movable(s state, from, to int) bool {
	if len(s[from]) == 0 {
		return false
	}
	if len(s[to]) == 0 {
		return true
	}
	if s[from][len(s[from])-1] > s[to][len(s[to])-1] {
		return true
	}
	return false
}

func reached(st state, n int) bool {
	return len(st[0]) == n || len(st[2]) == n
}
