package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (int, [][]int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	t := scanInt(sc)
	gains := make([][]int, t)
	for i := 0; i < t; i++ {
		gains[i] = make([]int, t+1)
		for j := 1; j <= t; j++ {
			gains[i][j] = scanInt(sc)
		}
	}
	return t, gains
}

func resolve(t int, gains [][]int) int {
	// dp[t] = max_{0<=i<j<t}( dp[i] + v[i][j] )
	// init
	dp := make([]int, t+2)
	for c := 1; c <= t+1; c++ {
		for j := c - 1; j > 0; j-- {
			for i := j - 1; i >= 0; i-- {
				changeToMax(&(dp[c]), dp[i]+gains[i][j])
			}
		}
	}
	return dp[t+1]
}

const (
	ninf = math.MinInt32
)

func changeToMin(v *int, cand int) (updated bool) {
	if *v > cand {
		*v = cand
		updated = true
	}
	return updated
}

func changeToMax(v *int, cand int) (updated bool) {
	if *v < cand {
		*v = cand
		updated = true
	}
	return updated
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
