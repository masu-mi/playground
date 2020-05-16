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
	states := make([]int, n+1)
	for i := 1; i <= n; i++ {
		states[i] = inf
	}
	for i := 1; i <= n; i++ {
		changeToMin(&(states[i]), states[i-1]+1)
		for j := 6; j <= i; j *= 6 {
			changeToMin(&(states[i]), states[i-j]+1)
		}
		for j := 9; j <= i; j *= 9 {
			changeToMin(&(states[i]), states[i-j]+1)
		}
	}
	return states[n]
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
