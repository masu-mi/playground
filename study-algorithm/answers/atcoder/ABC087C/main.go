package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	grid := make([][]int, 2)
	getCandies := make([][]int, 2)
	for i := 0; i < 2; i++ {
		grid[i] = make([]int, n)
		getCandies[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sc.Scan()
			grid[i][j], _ = strconv.Atoi(sc.Text())
			max := 0
			if i > 0 {
				if v := getCandies[i-1][j]; max < v {
					max = v
				}
			}
			if j > 0 {
				if v := getCandies[i][j-1]; max < v {
					max = v
				}
			}
			getCandies[i][j] = max + grid[i][j]
		}
	}
	fmt.Printf("%d\n", getCandies[1][n-1])
}
