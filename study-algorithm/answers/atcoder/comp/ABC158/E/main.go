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

func parseProblem(r io.Reader) (int, int, string) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	p := scanInt(sc)
	s := scanString(sc)

	return n, p, s
}

func resolve(n, p int, s string) int {
	upper = map[int]int{0: 1}
	if p == 2 || p == 5 {
		return countLastNumberIs(n, p, s)
	}
	count := 0
	nums := map[int]int{0: 1}
	l := len(s)
	before := 0
	for i := l; i > 0; i-- {
		current := toInt(s, i-1)

		upper[l-i+1] = (upper[l-i] * 10) % p
		current = (upper[l-i+1] * current) % p

		current = (before + current) % p
		count += nums[current]
		nums[current]++
		before = current
	}
	xc := 0
	for _, v := range nums {
		xc += v * (v + 1) / 2
	}
	return count
}

var upper map[int]int

func countLastNumberIs(n, p int, s string) int {
	num := 0
	for i := 0; i < n; i++ {
		v := toInt(s, i)
		if v%p == 0 {
			num += i + 1
		}
	}
	return num
}

func toInt(s string, i int) int {
	d, _ := strconv.Atoi(s[i : i+1])
	return d
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
