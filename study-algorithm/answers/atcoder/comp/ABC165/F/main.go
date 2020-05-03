package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	resolve(parseProblem(os.Stdin))
}

func parseProblem(r io.Reader) (int, []int, *graph) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = scanInt(sc)
	}
	g := scanGraph(n, n-1, sc)
	return n, nums, g
}

func resolve(n int, nums []int, g *graph) {
	s := status{
		nums:    nums,
		parents: make([]int, n),
		dp:      make([]int, n),
		line:    make([]int, n),
		ans:     make([]int, n),
	}
	for i := 0; i < n; i++ {
		s.line[i] = inf
	}
	s.parents[0] = -1
	s.dfs(0, g)
	for _, v := range s.ans {
		fmt.Println(v)
	}
}

type status struct {
	parents []int
	dp      []int
	line    []int
	nums    []int
	ans     []int
}

func (s *status) dfs(c int, g *graph) {
	idx := lowerBound(len(s.nums), func(i int) bool {
		return s.line[i] >= s.nums[c]
	})
	old := s.line[idx]
	s.line[idx] = s.nums[c]
	s.ans[c] = lowerBound(len(s.nums), func(i int) bool {
		return s.line[i] >= inf
	})
	defer func() {
		s.line[idx] = old
	}()
	for _, t := range g.edges[c].members() {
		if s.parents[c] == t {
			continue
		}
		s.parents[t] = c
		s.dfs(t, g)
	}
}

const (
	inf = math.MaxInt32
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

func lowerBound(n int, f func(i int) bool) int {
	return sort.Search(n, f)
}

func upperBoundOfLowSide(n int, f func(i int) bool) int {
	return sort.Search(n, func(i int) bool { return !f(i) }) - 1
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

func scanGraph(n, m int, sc *bufio.Scanner) *graph {
	g := newGraph(n)
	for i := 0; i < m; i++ {
		x, y := scanInt(sc), scanInt(sc)
		// 0-indexed
		x--
		y--
		g.addEdge(x, y)
	}
	return g
}

func (g *graph) addEdge(x, y int) {
	g.edges[x].add(y)
	g.edges[y].add(x)
}

func (g *graph) addDirectedEdge(x, y int) {
	g.edges[x].add(y)
}

func (g *graph) exists(x, y int) bool {
	return g.edges[x].doesContain(y)
}

type nodeSet map[int]none

func newSet() nodeSet {
	return make(map[int]none)
}

func (s nodeSet) add(item int) {
	s[item] = mark
}

func (s nodeSet) doesContain(item int) bool {
	_, ok := s[item]
	return ok
}

func (s nodeSet) size() int {
	return len(s)
}

func (s nodeSet) members() (l []int) {
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
