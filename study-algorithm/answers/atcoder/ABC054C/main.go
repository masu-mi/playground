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

type none struct{}

var mark none = struct{}{}

type edge struct{ x, y int }

type edgeSet map[edge]none

func (es edgeSet) exist(e edge) bool {
	x, y := min(e.x, e.y), max(e.x, e.y)
	_, ok := es[edge{x, y}]
	return ok
}

func (es edgeSet) Add(e edge) {
	x, y := min(e.x, e.y), max(e.x, e.y)
	es[edge{x, y}] = mark
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

func parseProblem(r io.Reader) (n, m int, es edgeSet) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	es = edgeSet{}

	n = scanInt(sc)
	m = scanInt(sc)
	for i := 0; i < m; i++ {
		x, y := scanInt(sc)-1, scanInt(sc)-1
		es.Add(edge{x, y})
	}
	return n, m, es
}

func resolve(n, m int, es edgeSet) int {
	num := 0
	for l := range permutations(n - 1) {
		if allExists(es, l) {
			num++
		}
	}
	return num
}

func allExists(es edgeSet, l []int) bool {
	cur := -1
	for i := 0; i < len(l); i++ {
		next := l[i]
		if !es.exist(edge{cur + 1, next + 1}) {
			return false
		}
		cur = next
	}
	return true
}

func permutations(l int) chan []int {
	ch := make(chan []int)
	go func() {
		dfsPermutations(0, make([]bool, l), []int{}, func(perm []int) bool {
			p := make([]int, len(perm))
			copy(p, perm)
			ch <- p
			return false
		})
		close(ch)
	}()
	return ch
}

func dfsPermutations(pos int, used []bool, perm []int, atLeaf func(perm []int) (halt bool)) (halt bool) {
	l := len(used)
	if pos == l {
		if atLeaf(perm) {
			return true
		}
	}

	for i := 0; i < l; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		if dfsPermutations(pos+1, used, append(perm, i), atLeaf) {
			return true
		}
		used[i] = false
	}
	return false
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
