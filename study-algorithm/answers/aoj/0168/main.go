package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	answer(os.Stdin)
}

func answer(r io.Reader) int {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	num := scanInt(sc)
	for ; num != 0; num = scanInt(sc) {
		fmt.Printf("%d\n", resolve(num))
	}
	return 0
}

func resolve(n int) int {
	states := make([]int, n+1)
	states[0] = 1
	for i := 0; i < n; i++ {
		// 配るDP
		for j := 1; j <= 3; j++ {
			if i+j <= n {
				states[i+j] += states[i]
			}
		}
	}
	num := states[n]
	year := (((num + 9) / 10) + 364) / 365
	return year
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
