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

func parseProblem(r io.Reader) (int, []int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	ps := []int{}
	for i := 0; i < n; i++ {
		ps = append(ps, scanInt(sc))
	}
	return n, ps
}

func resolve(n int, ps []int) int {
	dp := make([]map[int]struct{}, n+1)
	dp[0] = map[int]struct{}{0: struct{}{}}
	for i := 1; i <= n; i++ {
		dp[i] = map[int]struct{}{}
		for k := range dp[i-1] {
			dp[i][k] = struct{}{}
			dp[i][k+ps[i-1]] = struct{}{}
		}
	}
	return len(dp[n])
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
