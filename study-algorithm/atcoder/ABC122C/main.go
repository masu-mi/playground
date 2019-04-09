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
	num := 0
	for i := 1; i < n; i++ {
		if s[i] == 'C' && s[i-1] == 'A' {
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

// type acDFA struct{ s byte }
//
// func (m *acDFA) recognize(c byte) bool {
// 	if m.s == 'A' && c == 'C' {
// 		m.s = 'C'
// 		return true
// 	}
// 	m.s = c
// 	return false
// }

// func scanLine() (s string) {
// 	sc := bufio.NewScanner(os.Stdin)
// 	size := 100000
// 	sc.Buffer(make([]byte, size), size)
// 	sc.Scan()
// 	return sc.Text()
// }
