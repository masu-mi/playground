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

func parseProblem(r io.Reader) int {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	return scanInt(sc)
}

func resolve(n int) int {
	w := n
	dp := make([]int, w+1)
	var choices []int
	for i := 1; i <= w; i *= 6 {
		choices = append(choices, i)
	}
	for i := 9; i <= w; i *= 9 {
		choices = append(choices, i)
	}
	// init
	for i := 0; i <= w; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	for _, c := range choices {
		for i := c; i <= w; i++ {
			changeToMin(&(dp[i]), dp[i-c]+1)
		}
	}
	return dp[n]
}

const (
	inf  = math.MaxInt32
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
