package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	resolve(parseProblem(os.Stdin))
}

func parseProblem(r io.Reader) (int, int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, m := scanInt(sc), scanInt(sc)
	return n, m
}

func resolve(n, m int) {
	registered := make([]bool, n)
	count := 0
	idx := 0
	diff := 1
	for count < m && idx+diff < n {
		if registered[idx] {
			idx++
			continue
		}
		fmt.Printf("%d %d\n", idx+1, idx+1+diff)
		registered[idx+diff] = true
		count++
		diff++
		idx++
	}
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
