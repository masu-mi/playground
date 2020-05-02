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

func parseProblem(r io.Reader) (int, *graph, *bufio.Scanner) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	g := scanGraph(n, sc)
	return n, g, sc
}

func resolve(n int, g *graph, sc *bufio.Scanner) int {
	deg := make([]int, n)
	for k, s := range g.edges {
		deg[k] = s.size()
	}
	q := []int{}
	for k, d := range deg {
		if d == 1 { // leaf
			q = append(q, k)
		}
	}
	pushed := newSet()
	for len(q) > 0 {
		h := q[0]
		q = q[1:]
		pushed.add(node(h))
		deg[h]--
		for _, y := range g.edges[h].members() {
			deg[int(y)]--
			if deg[int(y)] == 1 {
				q = append(q, int(y))
			}
		}
	}
	m := scanInt(sc)
	for i := 0; i < m; i++ {
		x, y := scanInt(sc), scanInt(sc)
		x--
		y--
		if !pushed.doesContain(node(x)) && !pushed.doesContain(node(y)) {
			fmt.Println(2)
		} else {
			fmt.Println(1)
		}
	}
	return n
}

type graph struct {
	size  int
	edges []nodeSet
}

func newGraph(n int) *graph {
	g := &graph{
		size:  n,
		edges: make([]nodeSet, n),
	}
	for i := 0; i < n; i++ {
		g.edges[i] = newSet()
	}
	return g
}

func scanGraph(n int, sc *bufio.Scanner) *graph {
	g := newGraph(n)
	for i := 0; i < n; i++ {
		x, y := scanInt(sc), scanInt(sc)
		// 0-indexed
		x--
		y--
		g.addEdge(x, y)
	}
	return g
}

func (g *graph) addEdge(x, y int) {
	g.edges[x].add(node(y))
	g.edges[y].add(node(x))
}

func (g *graph) exists(x, y int) bool {
	return g.edges[x].doesContain(node(y))
}

type node int

type nodeSet map[node]none

func newSet() nodeSet {
	return make(map[node]none)
}

func (s nodeSet) add(item node) {
	s[item] = mark
}

func (s nodeSet) doesContain(item node) bool {
	_, ok := s[item]
	return ok
}

func (s nodeSet) size() int {
	return len(s)
}

func (s nodeSet) members() (l []node) {
	for k := range s {
		l = append(l, k)
	}
	return l
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
