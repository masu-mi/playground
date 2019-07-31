package main

import (
	"bytes"
	"errors"
	"fmt"
)

func main() {
	// discover articulation
	g := loadGraph()
	for a, _ := range findArticulations(g) {
		fmt.Printf("articulation: %d\n", a)
	}
}

func loadGraph() *Graph {
	var n, m int
	fmt.Scan(&n, &m)
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		g.AddEdge(u, v)
	}
	return g
}

func findArticulations(g *Graph) (articulations map[int]struct{}) {
	articulations = make(map[int]struct{})
	c := newContext(g.n)

	reachable := make(map[int]int, g.n)
	treeOutDegree := make(map[int]int, g.n)
	c.erly = func(c *context, u int) {
		reachable[u] = u
	}
	c.late = func(c *context, u int) {
		isRoot := c.parent[u] < 1
		if isRoot {
			if treeOutDegree[u] > 1 {
				articulations[u] = struct{}{}
			}
		} else {
			if reachable[u] == c.parent[u] {
				// parent
				if c.parent[c.parent[u]] > 0 { // parent isn't root
					articulations[c.parent[u]] = struct{}{}
				}
			} else if reachable[u] == u {
				// bridge
				if c.parent[c.parent[u]] > 0 { // parent isn't root
					articulations[c.parent[u]] = struct{}{}
				}
				if treeOutDegree[u] > 0 { // u have children
					articulations[u] = struct{}{}
				}
			}
		}
		uTime := c.entryTime[reachable[u]]
		pTime := c.entryTime[reachable[c.parent[u]]]
		if uTime < pTime {
			reachable[c.parent[u]] = reachable[u]
		}
	}
	c.processEdge = func(c *context, u, v int) {
		t := c.edgeType(u, v)
		switch t {
		case tree:
			treeOutDegree[u]++
		case back:
			if c.parent[u] != v && c.entryTime[v] < c.entryTime[u] {
				reachable[u] = v
			}
		default:
		}
	}
	time := 0
	c.dfs(g, 0, 1, &time)
	return articulations
}

type context struct {
	discovered  map[int]struct{}
	finished    map[int]struct{}
	entryTime   map[int]int
	parent      map[int]int
	erly, late  func(c *context, u int)
	processEdge func(c *context, u, v int)
}

func newContext(num int) *context {
	return &context{
		discovered: make(map[int]struct{}, num),
		finished:   make(map[int]struct{}, num),
		entryTime:  make(map[int]int, num),
		parent:     make(map[int]int, num),
	}
}

type EdgeType int

const (
	tree EdgeType = 0 + iota
	back
)

func (c *context) edgeType(u, v int) EdgeType {
	// this func assumes called in processEdge.
	if _, discovered := c.discovered[v]; !discovered {
		return tree
	}
	return back
}

func (c *context) dfs(g *Graph, p, u int, time *int) {
	c.parent[u] = p
	c.discovered[u] = struct{}{}
	*time += 1
	c.entryTime[u] = *time
	if c.erly != nil {
		c.erly(c, u)
	}
	for v := range g.edges[u] {
		if v == p {
			continue
		}
		if _, discovered := c.discovered[v]; !discovered {
			c.processEdge(c, u, v)
			c.dfs(g, u, v, time)
		} else if _, finished := c.finished[v]; !finished {
			c.processEdge(c, u, v)
		}
	}
	if c.late != nil {
		c.late(c, u)
	}
	*time += 1
	c.finished[u] = struct{}{}
}

type Graph struct {
	n     int // number of vertices; node id are required in [1..n]
	edges map[int]map[int]struct{}
}

func NewGraph(num int) *Graph {
	return &Graph{
		n:     num,
		edges: make(map[int]map[int]struct{}, num),
	}
}

func (g *Graph) ToString() string {
	displayed := map[struct{ u, v int }]struct{}{}
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprintln(buf, "graph test {")
	for u, e := range g.edges {
		for v := range e {
			if _, exists := displayed[struct{ u, v int }{u, v}]; !exists {
				fmt.Fprintln(buf, "%d -> %d", u, v)
				displayed[struct{ u, v int }{u, v}] = struct{}{}
				displayed[struct{ u, v int }{v, u}] = struct{}{}
			}
		}
	}
	fmt.Fprintln(buf, "}")
	return buf.String()
}

var (
	ErrorInvalidNode = errors.New("invalid Node")
	ErrorSelfLoop    = errors.New("self loop")
	ErrorMultiEdge   = errors.New("multiedge")
)

func (g *Graph) AddEdge(u, v int) error {
	if u <= 0 || u > g.n {
		return ErrorInvalidNode
	}
	if v <= 0 || v > g.n {
		return ErrorInvalidNode
	}
	if u == v {
		return ErrorSelfLoop
	}
	return g.addEdge(u, v)
}

func (g *Graph) addEdge(u, v int) error {
	ut, ok := g.edges[u]
	if !ok {
		ut = make(map[int]struct{})
	}
	if _, registered := ut[v]; registered {
		return ErrorMultiEdge
	}
	ut[v] = struct{}{}
	g.edges[u] = ut

	vt, ok := g.edges[v]
	if !ok {
		vt = make(map[int]struct{})
	}
	// pass to check multiedge, because of undirected graph's symmetry
	vt[u] = struct{}{}
	g.edges[v] = vt
	return nil
}
