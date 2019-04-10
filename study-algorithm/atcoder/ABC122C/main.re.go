package main

import (
	"fmt"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	var s string
	fmt.Scan(&s)
	culumtive := make([]int, n)
	num, m := 0, &dfa{}
	for i := 0; i < n; i++ {
		if m.recognize(s[i]) {
			num++
		}
		culumtive[i] = num
	}
	var l, r int
	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r)
		a := culumtive[r-1] - culumtive[l-1]
		fmt.Printf("%d\n", a)
	}
}

type dfa struct{ s byte }

func (m *dfa) recognize(c byte) bool {
	if m.s == 'A' && c == 'C' {
		m.s = 'C'
		return true
	}
	m.s = c
	return false
}
