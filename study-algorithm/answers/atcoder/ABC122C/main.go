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
