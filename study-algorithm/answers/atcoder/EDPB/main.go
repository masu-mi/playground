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

func parseProblem(r io.Reader) (int, int, []int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, k := scanInt(sc), scanInt(sc)
	// 0-indexed
	hs := make([]int, n)
	for i := 0; i < n; i++ {
		hs[i] = scanInt(sc)
	}
	return n, k, hs
}

func resolve(n, k int, hs []int) int {
	states := make([]int, n)
	states[0] = 0
	for i := 1; i < n; i++ {
		states[i] = inf
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			if i+j < n {
				changeToMin(&(states[i+j]), states[i]+abs(hs[i]-hs[i+j]))
			}
		}
	}
	return states[n-1]
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
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
