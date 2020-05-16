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

func parseProblem(r io.Reader) (string, string) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	return scanString(sc), scanString(sc)
}

func resolve(s, t string) int {
	// init
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	// dp[i][j] := max length of LCS of s[:i], t[:j]
	// dp[i][j] := max(dp[i-1][j-1] +1/0, dp[i-1][j] + 1/0, dp[i][j-1] + 1/0)
	type diff struct{ di, dj int }
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			changeToMax(&(dp[i][j]), dp[i][j-1])
			changeToMax(&(dp[i][j]), dp[i-1][j])
			if s[i-1] == t[j-1] {
				changeToMax(&(dp[i][j]), dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(s)][len(t)]
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
