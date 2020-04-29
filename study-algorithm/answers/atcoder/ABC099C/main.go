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
	// index == value
	states := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		states[i] = inf
	}
	// 配るDP
	for i := 0; i < n; i++ {
		if i+1 <= n {
			changeToMin(&(states[i+1]), states[i]+1)
		}
		for idx := 6; i+idx <= n; idx *= 6 {
			changeToMin(&(states[i+idx]), states[i]+1)
		}
		for idx := 9; i+idx <= n; idx *= 9 {
			changeToMin(&(states[i+idx]), states[i]+1)
		}
	}
	return states[n]
}

const (
	inf = math.MaxInt32
)

func changeToMin(v *int, cand int) (updated bool) {
	if *v > cand {
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
