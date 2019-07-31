package main

import (
	"fmt"
)

func main() {
	g := loadGraph()
	for a := range findArticulation(g) {
		fmt.Printf("articulation: %d\n", a)
	}
}

type none struct{}

var mark none = struct{}{}

func findArticulation(g *graph) (articulations map[int]none) {
	// write code as dfs
	s := newDfsState(g)
	s.findArticulation(0, 1, 0)
	return s.articulations
}

type dfsState struct {
	// lowlink
	ord map[int]int
	low map[int]int
	// result
	articulations map[int]none

	// standard dfs
	g          *graph
	discovered map[int]none
	finished   map[int]none
	parent     map[int]int
	childNum   map[int]int
}

func newDfsState(g *graph) *dfsState {
	return &dfsState{
		g:          g,
		ord:        make(map[int]int),
		low:        make(map[int]int),
		discovered: make(map[int]none),
		finished:   make(map[int]none),
		parent:     make(map[int]int),
		childNum:   make(map[int]int),

		articulations: make(map[int]none),
	}
}

func (s *dfsState) findArticulation(p, u int, ord int) int {
	ord++
	s.parent[u] = p
	s.discovered[u] = mark
	s.ord[u] = ord
	s.low[u] = ord
	for _, v := range s.g.edges[u] {
		if v == p {
			continue
		}
		switch s.getEdgeType(u, v) {
		case tree:
			s.childNum[u]++
			ord = s.findArticulation(u, v, ord)
		default:
			// back
			if s.ord[v] < s.low[u] {
				s.low[u] = s.ord[v]
			}
		}
	}
	// late phase
	// set parent's low
	uLow := s.low[u]
	pLow := s.low[s.parent[u]]
	if uLow < pLow {
		s.low[s.parent[u]] = uLow
	}
	// judge articulation
	if isRoot := s.parent[u] < 1; isRoot {
		if s.childNum[u] > 1 {
			s.articulations[u] = mark
		}
	} else if s.low[u] >= s.ord[u] && s.childNum[u] > 1 {
		s.articulations[u] = mark
	}
	s.finished[u] = mark
	return ord + 1
}

type edgeType int

const (
	tree edgeType = 0 + iota
	back
)

func (s dfsState) getEdgeType(u, v int) edgeType {
	if _, ok := s.discovered[v]; !ok {
		return tree
	}
	return back
}

func loadGraph() *graph {
	var n, m int
	fmt.Scan(&n, &m)
	g := newGraph(n)
	var u, v int
	for i := 0; i < m; i++ {
		fmt.Scan(&u, &v)
		g.addEdge(u, v)
	}
	return g
}

type graph struct {
	num   int
	edges map[int][]int
}

func newGraph(num int) *graph {
	return &graph{
		num:   num,
		edges: make(map[int][]int, num),
	}
}

func (g *graph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
	g.edges[v] = append(g.edges[v], u)
}
