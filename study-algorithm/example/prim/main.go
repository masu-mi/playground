package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	graph := parseProblem(os.Stdin)
	start, goal := 0, graph.nodeNum-1
	fmt.Println("digraph g {")
	for c, p := range getMinimumSpanningTree(graph, start, goal) {
		fmt.Printf("\t%d -> %d;\n", p, c)
	}
	fmt.Println("}")
}

type mark struct{}

var note = mark{}

func getMinimumSpanningTree(g *graph, s, t int) map[int]int {
	// WARGIN this code is example.
	// This function assumes g is connected graph.
	// implement with Prim
	parent := map[int]int{0: 0}
	heap := newHeap()
	current := 0
	for len(parent) < g.nodeNum {
		for _, edge := range g.edges[current] {
			if _, ok := parent[edge.to]; !ok {
				heap.push(part{edge: edge})
			}
		}
		p := part{edge: edge{to: -1}}
		for !heap.isEmpty() {
			p, _ = heap.pop()
			if _, ok := parent[p.edge.to]; !ok {
				break
			}
		}
		if p.edge.to == -1 {
			fmt.Println("ERROR")
			// ERROR
			break
		}
		parent[p.edge.to] = p.edge.from
		current = p.edge.to
	}
	return parent
}

type heap struct {
	length int
	values []part
}

func newHeap() *heap {
	return &heap{}
}
func (h *heap) isEmpty() bool {
	return h.length == 0
}
func (h *heap) push(p part) {
	h.length++
	if len(h.values) < h.length {
		h.values = append(h.values, p)
	} else {
		h.values[h.length-1] = p
	}
	h.recoverPushedRank()
}
func (h *heap) pop() (part, error) {
	if h.length == 0 {
		return part{edge: edge{to: -1}}, errors.New("empty heap")
	}
	r := h.values[0]
	h.values[0] = h.values[h.length-1]
	h.length--
	h.recoverPopedRank()
	return r, nil
}

func (h *heap) recoverPopedRank() {
	c := 0
	for c < h.length {
		cCost := h.values[c].cost
		rIdx := 2 * (c + 1)
		lIdx := 2*(c+1) + 1
		if rIdx < h.length {
			rCost := h.values[rIdx].cost
			if lIdx < h.length {
				lCost := h.values[lIdx].cost
				if cCost > rCost && cCost > lCost {
					if rCost < lCost {
						h.values[rIdx], h.values[c] = h.values[c], h.values[rIdx]
						c = rIdx
					} else {
						h.values[lIdx], h.values[c] = h.values[c], h.values[lIdx]
						c = lIdx
					}
				}
				if cCost > rCost {
					h.values[rIdx], h.values[c] = h.values[c], h.values[rIdx]
					c = rIdx
				} else {
					break
				}
			} else {
				h.values[rIdx], h.values[c] = h.values[c], h.values[rIdx]
				c = rIdx
			}
		} else if lIdx < h.length {
			lCost := h.values[lIdx].cost
			if cCost > lCost {
				h.values[lIdx], h.values[c] = h.values[c], h.values[lIdx]
				c = lIdx
			} else {
				break
			}
		} else {
			break
		}
	}
}

func (h *heap) recoverPushedRank() {
	c := h.length - 1
	for c > 0 {
		if h.values[c].cost < h.values[(c+1)/2-1].cost {
			h.values[(c+1)/2-1], h.values[c] = h.values[c], h.values[(c+1)/2-1]
			c = (c+1)/2 - 1
		} else {
			break
		}
	}
}

type part struct {
	cost int
	edge
}

type edge struct {
	from, to, weight int
}
type graph struct {
	nodeNum int
	edges   map[int][]edge
}

func (g *graph) addEdge(u, v, w int) {
	g.edges[u] = append(g.edges[u], edge{from: u, to: v, weight: w})
	g.edges[v] = append(g.edges[v], edge{from: v, to: u, weight: w})
}

func parseProblem(r io.Reader) *graph {
	var n, m int
	fmt.Fscan(r, &n)
	g := &graph{nodeNum: n, edges: make(map[int][]edge)}

	fmt.Fscan(r, &m)
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
