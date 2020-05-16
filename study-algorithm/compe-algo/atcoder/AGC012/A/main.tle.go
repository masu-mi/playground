package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, 3*n)
	for k := range as {
		fmt.Scan(&as[k])
	}
	sort.Ints(as)
	var sum int
	for i := n; i < 3*n; i += 2 {
		sum += as[i]
	}
	fmt.Printf("%d\n", sum)
}
