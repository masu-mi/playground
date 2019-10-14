package main

import (
	"fmt"
	"io"
	"os"

	"github.com/k0kubun/pp"
)

func main() {
	big := parseInput(os.Stdin)
	fmt.Println("[Problem] Input BiGraph")
	pp.Println(big)
	result := searchMaximumMatching(big)
	fmt.Println("[Result] MaximumMatching")
	for _, e := range result {
		fmt.Printf("%d - %d\n", e.f, e.s)
	}
}

type edge struct{ f, s int }

func searchMaximumMatching(g *bigraph) []edge {
	registered := make(map[edge]marker)
	arrived := make(map[int]marker)
	for fid := range g.firstGroups {
		// dummy edge {0, 0} for preventing panic at slicing in after()
		paths := []edge{edge{0, 0}}
		c := newContext(
			func(p, c id) (stop bool) {
				var candidate edge
				if p.cat == 1 {
					candidate = edge{p.num, c.num}
				} else {
					candidate = edge{c.num, p.num}
				}

				paths = append(paths, candidate)
				if _, ok := registered[candidate]; !ok {
					if _, ok := arrived[candidate.s]; ok {
						return false
					}
					opIsDelete := false
					for i := len(paths); i > 0; i-- {
						if opIsDelete {
							delete(registered, paths[i-1])
							opIsDelete = false
						} else {
							registered[paths[i-1]] = mark
							arrived[paths[i-1].s] = mark
							opIsDelete = true
						}
					}
					return true
				}
				return false
			},
			func() {
				paths = paths[0 : len(paths)-1]
			},
		)
		startAt := 0
		c.dfs(g.graph, nullNode, fid, &startAt)
	}
	var result []edge
	for e := range registered {
		result = append(result, e)
	}
	return result
}

func parseInput(r io.Reader) *bigraph {
	var e int
	fmt.Fscan(r, &e)
	g := newBiGraph()

	for i := 0; i < e; i++ {
		var f, s int
		fmt.Fscan(r, &f, &s)
		g.addEdge(f, s)
	}
	return g
}

type marker struct{}

var mark = marker{}

type id struct{ cat, num int }

// We manage nodes' group (only 2 groups)
type bigraph struct {
	*graph
	firstGroups map[id]marker
}

func newBiGraph() *bigraph {
	return &bigraph{
		graph:       newGraph(),
		firstGroups: make(map[id]marker),
	}
}

func (bi *bigraph) addEdge(f, s int) {
	bi.firstGroups[id{1, f}] = mark
	bi.graph.addEdge(id{1, f}, id{2, s})
}

// brief linkedlist for graph
type graph struct {
	ids   map[id]marker
	edges map[id][]id
}

func newGraph() *graph {
	return &graph{
		ids:   make(map[id]marker),
		edges: make(map[id][]id),
	}
}

func (g *graph) addEdge(f, s id) {
	g.ids[f] = mark
	g.ids[s] = mark
	g.edges[f] = append(g.edges[f], s)
	g.edges[s] = append(g.edges[s], f)
}

type context struct {
	discovered map[id]marker
	finished   map[id]marker
	entryTime  map[id]int
	parent     map[id]id
	f          func(p, c id) (stop bool)
	after      func()
}

func newContext(f func(p, c id) (stop bool), after func()) *context {
	return &context{
		discovered: make(map[id]marker),
		finished:   make(map[id]marker),
		entryTime:  make(map[id]int),
		parent:     make(map[id]id),
		f:          f,
		after:      after,
	}
}

type edgeType int

const (
	back edgeType = 0 + iota
	tree
)

var nullNode = id{0, 0}

func (c *context) edgeType(f, t id) edgeType {
	if _, ok := c.discovered[t]; ok {
		return back
	}
	return tree
}

func (c *context) dfs(g *graph, p, ci id, time *int) bool {
	c.parent[ci] = p
	c.discovered[ci] = mark
	*time++
	c.entryTime[ci] = *time
	for _, cc := range g.edges[ci] {
		if cc == p {
			continue
		}
		// ignore back edge
		if c.edgeType(ci, cc) == back {
			continue
		}
		if c.f(ci, cc) {
			return true
		}
		if c.dfs(g, ci, cc, time) {
			return true
		}
	}
	c.after()
	*time++
	c.finished[ci] = mark
	return false
}
