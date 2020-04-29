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
	q := []int{}
	states := make([]int, n+1)
	for i := 0; i <= n; i++ {
		states[i] = -1
	}
	states[0] = 0
	q = append(q, 0)
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		if c == n {
			break
		}
		for i := 1; c+i <= n; i *= 6 {
			if states[c+i] == -1 {
				states[c+i] = states[c] + 1
				q = append(q, c+i)
			}
		}
		for i := 1; c+i <= n; i *= 9 {
			if states[c+i] == -1 {
				states[c+i] = states[c] + 1
				q = append(q, c+i)
			}
		}
	}
	return states[n]
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
