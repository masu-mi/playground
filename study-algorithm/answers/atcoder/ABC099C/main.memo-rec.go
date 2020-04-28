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

	n := scanInt(sc)
	states = make([]int, n+1)
	for i := 0; i <= n; i++ {
		states[i] = -1
	}
	return n
}

var states []int

func resolve(n int) (v int) {
	defer func() {
		states[n] = v
	}()
	if n == 0 {
		v = 0
		return
	}
	if v := states[n]; v != -1 {
		return v
	}

	v = n
	for i := 1; i <= n; i *= 6 {
		changeToMin(&v, resolve(n-i)+1)
	}
	for i := 1; i <= n; i *= 9 {
		changeToMin(&v, resolve(n-i)+1)
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
