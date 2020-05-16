package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	min := math.MaxInt32
	for i := 1; i <= n/2; i++ {
		tmp := sumValues(i, n-i)
		if tmp < min {
			min = tmp
		}
	}
	fmt.Printf("%d\n", min)
}

func sumValues(x, y int) int {
	return sum(x) + sum(y)
}
func sum(i int) (r int) {
	tmp := i
	for tmp != 0 {
		r += tmp % 10
		tmp /= 10
	}
	return r
}
