package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	s := scanLine()
	culumtive := make([]int, n)
	m, num := &acDFA{}, 0
	for i := 0; i < n; i++ {
		if m.recognize(s[i]) {
			num++
		}
		culumtive[i] = num
	}
	var l, r int
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&l)
		fmt.Scan(&r)
		ans[i] = culumtive[r-1] - culumtive[l-1]
	}
	for _, a := range ans {
		fmt.Printf("%d\n", a)
	}
}

type state int
type acDFA struct {
	s state
}

func (m *acDFA) recognize(c byte) bool {
	if c == 'A' {
		m.s = 1
		return false
	}
	if m.s == 1 && c == 'C' {
		return true
	}
	m.s = 0
	return false
}

func scanLine() (s string) {
	sc := bufio.NewScanner(os.Stdin)
	size := 100000
	sc.Buffer(make([]byte, size), size)
	sc.Scan()
	return sc.Text()
}
