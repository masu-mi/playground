package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	ts := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ts[i])
	}
	fmt.Printf("%d\n", minTime(ts))
}

func minTime(ts []int) int {
	// loop
	min := math.MaxInt32
	for i := 0; i < 1<<uint(len(ts)); i++ {
		lt, rt, idx := 0, 0, 0
		for j := 1; j < 1<<uint(len(ts)); j <<= 1 {
			if i&j == 0 {
				lt += ts[idx]
			} else {
				rt += ts[idx]
			}
			idx++
		}
		v := lt
		if lt < rt {
			v = rt
		}
		if v < min {
			min = v
		}
	}
	return min
}
