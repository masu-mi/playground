package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Printf("%d\n", overwrapLength(a, b, c, d))
}

func overwrapLength(a, b, c, d int) int {
	min := math.MaxInt32
	for _, l := range []int{d - a, b - c, b - a, d - c} {
		if l < min {
			min = l
		}
	}
	if min < 0 {
		return 0
	}
	return min
}
