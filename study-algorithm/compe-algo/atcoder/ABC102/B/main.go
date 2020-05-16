package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var (
		min = +math.MaxInt32
		max = -math.MaxInt32
	)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		if tmp < min {
			min = tmp
		}
		if tmp > max {
			max = tmp
		}
	}
	fmt.Printf("%d\n", max-min)
}
