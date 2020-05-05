package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (int, []int, []int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	hs := make([]int, n)
	as := make([]int, n)
	for i := 0; i < n; i++ {
		hs[i] = scanInt(sc)
	}
	for i := 0; i < n; i++ {
		as[i] = scanInt(sc)
	}
	return n, hs, as
}

func resolve(n int, hs, as []int) int {
	// dp min hight dp[i][j] items < i, hight = j
	// dp[i+1][h] = max_(k<h)(dp[i][k] + a[i]) // h[i] == h
	// dp[i+1][h] = dp[i][h]                   // h[i] != h
	// -> dp[h[i]] = max_(k<h[i]){dp[k]+a[i]}
	l := 0
	for i := 0; i < len(hs); i++ {
		if l < hs[i] {
			l = hs[i]
		}
	}
	dp := make([]int, l+1)
	for i := 0; i < n; i++ {
		vs := []int{}
		for j := 0; j < hs[i] && j <= l; j++ {
			vs = append(vs, dp[j]+as[i])
		}
		dp[hs[i]] = max(vs)
	}
	result := 0
	for i := 0; i < len(dp); i++ {
		if v := dp[i]; result < v {
			result = v
		}
	}
	return result
}

func max(a []int) int {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
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
