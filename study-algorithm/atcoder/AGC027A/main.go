package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}
	sort.Sort(sort.IntSlice(as))
	var distributed, i int
	for true {
		if i == len(as) {
			i--
			break
		}
		if distributed+as[i] > x {
			break
		}
		distributed += as[i]
		i++
		if distributed == x {
			break
		}
	}
	fmt.Printf("%d\n", i)
}
