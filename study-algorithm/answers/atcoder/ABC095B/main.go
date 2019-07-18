package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	min := math.MaxInt32
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		x -= tmp
		if tmp < min {
			min = tmp
		}
	}
	fmt.Printf("%d\n", n+x/min)
}
