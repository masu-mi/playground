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

func parseProblem(r io.Reader) (int, int, int, []int, []int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, w := scanInt(sc), scanInt(sc)
	ws, vs := make([]int, n), make([]int, n)
	maxValue := 0
	for i := 0; i < n; i++ {
		ws[i] = scanInt(sc)
		vs[i] = scanInt(sc)
		maxValue += vs[i]
	}
	return n, w, maxValue, ws, vs
}

func resolve(n, w, maxValue int, ws, vs []int) int {
	// keep min weight
	states := make([][]int, n+1)
	states[0] = make([]int, maxValue+1)
	for i := 0; i <= maxValue; i++ {
		states[0][i] = inf
	}
	states[0][0] = 0
	for i := 1; i <= n; i++ {
		states[i] = make([]int, maxValue+1)
		for j := 0; j <= maxValue; j++ {
			states[i][j] = inf
		}
		for v := 0; v <= maxValue; v++ {
			if v-vs[i-1] >= 0 {
				changeToMin(&(states[i][v]), states[i-1][v-vs[i-1]]+ws[i-1])
			}
			changeToMin(&(states[i][v]), states[i-1][v])
		}
	}
	result := 0
	for i := 0; i <= maxValue; i++ {
		if states[n][i] <= w {
			result = i
		}
	}
	return result
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
