package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	parseProblem(os.Stdin)
}

func parseProblem(r io.Reader) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, k, m := scanInt(sc), scanInt(sc), scanInt(sc)
	target := n * m
	sum := 0
	for i := 0; i < n-1; i++ {
		sum += scanInt(sc)
	}
	lastTarget := target - sum
	if lastTarget < 0 {
		lastTarget = 0
	} else if lastTarget > k {
		lastTarget = -1
	}
	fmt.Println(lastTarget)
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
