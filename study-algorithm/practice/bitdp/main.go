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

func parseProblem(r io.Reader) (n, m int, es edgeSet) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, m = scanInt(sc), scanInt(sc)
	es = newSet()
	for i := 0; i < m; i++ {
		x, y, c := scanInt(sc), scanInt(sc), scanInt(sc)
		es.add(edge{x, y}, c)
	}
	return n, m, es
}

func resolve(n, m int, es edgeSet) int {
	states := make([][]int, n)
	prev := make([][]int, n)
	for i := range states {
		states[i] = baseStates(n)
		prev[i] = baseStates(n)
	}
	// init
	for i := 0; i < n; i++ {
		states[i][1<<uint(i)] = 0
		prev[i][1<<uint(i)] = -1
	}
	for s := 1; s < 1<<uint(n); s++ {
		for j := 0; j < n; j++ {
			if s&(1<<uint(j)) == 0 {
				continue
			}
			for k := 0; k < n; k++ {
				if s&(1<<uint(k)) != 0 {
					continue
				}
				baseCost := states[j][s]
				edgeCost := es.cost(edge{j, k})
				cost := baseCost + edgeCost
				if baseCost == math.MaxInt64 || edgeCost == math.MaxInt64 {
					cost = math.MaxInt64
				}
				// fmt.Printf("j: %d, k: %d(b%05b)\n", j, k, s|1<<uint(k))
				if states[k][s|1<<uint(k)] > cost {
					states[k][s|(1<<uint(k))] = cost
					prev[k][s|(1<<uint(k))] = j
				}
			}
		}
	}
	result := math.MaxInt64
	last := -1
	for i := range states {
		if v := states[i][1<<uint(n)-1]; result > v {
			result = v
			last = i
		}
	}
	s := 1<<uint(n) - 1
	for last != -1 {
		fmt.Printf("[node: %d], ", last)
		last, s = prev[last][s], s&(^(1 << uint(last)))
	}
	fmt.Println("")
	return result
}

func baseStates(n int) []int {
	states := make([]int, 1<<uint(n)+1)
	for i := 0; i < 1<<uint(n)+1; i++ {
		states[i] = math.MaxInt64
	}
	return states
}

type edge struct{ s, t int }

type edgeSet map[edge]int

func newSet() edgeSet {
	return make(map[edge]int)
}

func (s edgeSet) add(item edge, c int) {
	canonicalized := edge{min(item.s, item.t), max(item.s, item.t)}
	s[canonicalized] = c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s edgeSet) cost(item edge) int {
	canonicalized := edge{min(item.s, item.t), max(item.s, item.t)}
	c, ok := s[canonicalized]
	if !ok {
		return math.MaxInt64
	}
	return c
}

func (s edgeSet) size() int {
	return len(s)
}

var mark none

type none struct{}

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
