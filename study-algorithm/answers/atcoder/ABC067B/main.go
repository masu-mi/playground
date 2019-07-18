package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	ls := make([]int, n)
	for k := range ls {
		fmt.Scan(&ls[k])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ls)))
	var sum int
	for i := 0; i < k; i++ {
		sum += ls[i]
	}
	fmt.Printf("%d\n", sum)
}
