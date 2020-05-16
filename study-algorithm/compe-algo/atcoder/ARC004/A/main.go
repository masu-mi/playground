package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	xs := make([]int, n)
	ys := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&xs[i], &ys[i])
	}
	var max float64
	for i := 0; i < n; i++ {
		for k := i + 1; k < n; k++ {
			d := math.Sqrt(
				math.Pow(float64(xs[k]-xs[i]), 2.0) + math.Pow(float64(ys[k]-ys[i]), 2.0),
			)
			if d > max {
				max = d
			}
		}
	}
	fmt.Printf("%f\n", max)
}
