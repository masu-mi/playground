package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	max, min, tmp := -1, 1001, 0
	for i := n; i > 0; i-- {
		fmt.Scan(&tmp)
		if tmp > max {
			max = tmp
		}
		if tmp < min {
			min = tmp
		}
	}
	fmt.Printf("%d\n", max-min)
}
