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
	offset := 1
	startOffset := m + 1
	for i := m; i > 0; i-- {
		base := offset
		if (m-i)%2 == 1 {
			base += startOffset
		}
		fmt.Printf("%d %d\n", base, base+i)
		if (m-i)%2 == 1 {
			offset++
		}
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
