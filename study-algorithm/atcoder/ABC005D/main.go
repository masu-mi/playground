package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sums := make([][]int, n+1)
	sums[0] = make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			var a int
			fmt.Scan(&a)
			sums[i+1][j+1] = sums[i+1][j] + sums[i][j+1] - sums[i][j] + a
		}
	}
	maxScore := make([]int, n*n+1)
	for li := 0; li < n; li++ {
		for ri := li + 1; ri < n+1; ri++ {
			for lj := 0; lj < n; lj++ {
				for rj := lj + 1; rj < n+1; rj++ {
					size := (ri - li) * (rj - lj)
					score := sums[ri][rj] - sums[ri][lj] - sums[li][rj] + sums[li][lj]
					if maxScore[size] < score {
						maxScore[size] = score
					}
				}
			}
		}
	}
	for i := 1; i <= n*n; i++ {
		if p := maxScore[i-1]; maxScore[i] < p {
			maxScore[i] = p
		}
	}
	var q, p int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&p)
		fmt.Printf("%d\n", maxScore[p])
	}
}
