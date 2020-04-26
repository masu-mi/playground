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
	hs := make([]int, n+3)
	for i := 0; i < n; i++ {
		hs[i] = scanInt(sc)
	}
	return n, hs
}

func resolve(n int, hs []int) int {
	states := make([]int, n+3)
	for i := 0; i < len(states); i++ {
		states[i] = math.MaxInt32
	}
	states[0] = 0
	for i := 0; i < n; i++ {
		changeToMin(&(states[i+1]), states[i]+abs(hs[i]-hs[i+1]))
		changeToMin(&(states[i+2]), states[i]+abs(hs[i]-hs[i+2]))
	}
	return states[n-1]
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

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
