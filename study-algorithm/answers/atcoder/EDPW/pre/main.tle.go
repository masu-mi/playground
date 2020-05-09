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

func parseProblem(r io.Reader) (int, map[int][]query) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, m := scanInt(sc), scanInt(sc)
	qs := make(map[int][]query, m)
	for i := 0; i < m; i++ {
		l, r, a := scanInt(sc), scanInt(sc)+1, scanInt(sc)
		qs[r] = append(qs[r], query{l: l, r: r, a: a})
	}
	return n, qs
}

type query struct {
	l, r, a int
}

func resolve(n int, qs map[int][]query) int {
	dp := make([]int, n+2)
	for i := 1; i <= n+1; i++ {
		max := 0
		for k := 0; k < i; k++ {
			if max < dp[k] {
				max = dp[k]
			}
		}
		sum := 0
		for _, v := range qs[i] {
			sum += v.a
		}
		dp[i] = max + sum
		for j := i - 1; j > 0; j-- {
			for _, v := range qs[i] {
				if v.l < j {
					dp[j] += v.a
				}
			}
		}
	}
	max := 0
	for i := 0; i <= n+1; i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
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
