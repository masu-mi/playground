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

func parseProblem(r io.Reader) (int, int, int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	a, b, k := scanInt(sc), scanInt(sc), scanInt(sc)
	return a, b, k
}

func resolve(a, b, k int) int {
	max := min(a, b)
	count := 0
	for i := max; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			count++
		}
		if count == k {
			return i
		}
	}
	// out of scope
	return -1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
