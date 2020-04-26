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

func parseProblem(r io.Reader) (int, []map[byte]int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	n := scanInt(sc)
	// 0-indexed
	selections := []map[byte]int{}
	for i := 0; i < n; i++ {
		selections = append(selections, map[byte]int{
			'a': scanInt(sc),
			'b': scanInt(sc),
			'c': scanInt(sc),
		})
	}

	return n, selections
}

func resolve(n int, selections []map[byte]int) int {
	// 0-indexed
	states := make([]map[byte]int, n)
	// init
	for i := 0; i < n; i++ {
		states[i] = map[byte]int{
			'a': ninf,
			'b': ninf,
			'c': ninf,
		}
	}
	for _, c := range []byte{'a', 'b', 'c'} {
		states[0][c] = selections[0][c]
	}
	// 貰うDP
	for i := 1; i < n; i++ {
		for _, c := range []byte{'a', 'b', 'c'} {
			bC := states[i-1][c]
			for _, k := range []byte{'a', 'b', 'c'} {
				if c == k {
					continue
				}
				if states[i][k] < bC+selections[i][k] {
					states[i][k] = bC + selections[i][k]
				}
			}
		}
	}
	result := ninf
	for _, k := range []byte{'a', 'b', 'c'} {
		changeToMax(&result, states[n-1][k])
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
