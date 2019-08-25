package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	g := parseProblem(os.Stdin)
	printTree(createMinimumSpanningTree(g))
}

func printTree(t map[int][]int) {
	fmt.Println("graph g {")
	for c, edges := range t {
		for _, p := range edges {
			fmt.Printf("\t%d -- %d;\n", p, c)
		}
	}
	fmt.Println("}")
}

func createMinimumSpanningTree(g *graph) map[int][]int {
	var eds []edge
	for _, list := range g.edges {
		for _, e := range list {
			eds = append(eds, e)
		}
	}
	// It's more efficient to implement with heap.
	// But these edges are used with statistic order and sort is more easy to implement without heap package.
	result := map[int][]int{}
	uf := newUnionFind(g.nodes)
	sort.Sort(edges(eds))
	num := 0
	for _, e := range eds {
		if num == g.nodes-1 {
			break
		}
		f, t := e.f, e.t
		if uf.find(f) != uf.find(t) {
			result[t] = append(result[t], f)
			result[f] = append(result[f], t)
			uf.union(f, t)
			num++
		}
	}
	return result
}

type edges []edge

func (es edges) Len() int {
	return len([]edge(es))
}
func (es edges) Less(i, j int) bool {
	return es[i].w < es[j].w
}
func (es edges) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

type graph struct {
	nodes int
	edges map[int][]edge
}

type edge struct{ f, t, w int }

func newGraph(n int) *graph {
	return &graph{nodes: n, edges: make(map[int][]edge)}
}

func (g *graph) addEdge(u, v, w int) error {
	if u >= g.nodes || v >= g.nodes {
		return errors.New("invalid node")
	}
	g.edges[u] = append(g.edges[u], edge{f: u, t: v, w: w})
	g.edges[v] = append(g.edges[v], edge{f: v, t: u, w: w})
	return nil
}

type unionFind struct {
	length int
	parent []int
	size   []int
}

func newUnionFind(num int) *unionFind {
	u := &unionFind{
		length: num,
		parent: make([]int, num),
		size:   make([]int, num),
	}
	for i := 0; i < num; i++ {
		u.parent[i] = i
		u.size[i] = 1
	}
	return u
}

func (uf *unionFind) find(n int) int {
	p := uf.parent[n]
	if p == n {
		return p
	}
	return uf.find(p)
}

func (uf *unionFind) union(u, v int) {
	ru, rv := uf.find(u), uf.find(v)
	if ru == rv {
		return
	}
	if uf.size[ru] < uf.size[rv] {
		uf.size[rv] = uf.size[rv] + uf.size[ru]
		uf.parent[ru] = rv
	} else {
		uf.size[ru] = uf.size[ru] + uf.size[rv]
		uf.parent[rv] = ru
	}
}

func parseProblem(r io.Reader) *graph {
	var n, m int
	fmt.Scan(&n, &m)
	g := newGraph(n)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		w, _ := strconv.Atoi(sc.Text())
		g.addEdge(u, v, w)
	}
	return g
}
