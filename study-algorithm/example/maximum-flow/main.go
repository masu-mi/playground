package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	cg := parseProblem(os.Stdin)
	fmt.Println("[PROBLEM]")
	for _, es := range cg.edges {
		for _, e := range es {
			fmt.Printf("    %d -[%d]-> %d\n", e.f, e.c, e.t)
		}
	}
	fmt.Println("[Maximu flow]")
	for _, f := range findMaxFlow(cg) {
		fmt.Printf("    %d -[%d]-> %d\n", f.f, f.v, f.t)
	}
}

func parseProblem(r io.Reader) *capacityGraph {
	var n, m int
	fmt.Fscan(r, &n, &m)
	cg := newCapacityGraph(n)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		sc.Scan()
		f, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		c, _ := strconv.Atoi(sc.Text())
		cg.AddEdge(f, t, c)
	}
	return cg
}

var (
	// ErrInvalidNode means invalid node
	ErrInvalidNode = errors.New("invalid node")
	// ErrInvalidCapacity means invalid capacity edge
	ErrInvalidCapacity = errors.New("invalid capacity")
)

type edge struct{ f, t, c int }

type capacityGraph struct {
	nodeNum int
	edges   map[int]map[int]edge
}

func newCapacityGraph(n int) *capacityGraph {
	return &capacityGraph{
		nodeNum: n,
		edges:   make(map[int]map[int]edge),
	}
}

func (cg *capacityGraph) AddEdge(f, t, c int) error {
	if !cg.isValidNode(f) {
		return ErrInvalidNode
	}
	if !cg.isValidNode(t) {
		return ErrInvalidNode
	}
	if c < 1 {
		return ErrInvalidCapacity
	}
	cg.addEdge(f, t, c)
	return nil
}

func (cg *capacityGraph) isValidNode(n int) bool {
	if n < 1 || n > cg.nodeNum {
		return false
	}
	return true
}

func (cg *capacityGraph) addEdge(f, t, c int) {
	if _, ok := cg.edges[f]; !ok {
		cg.edges[f] = make(map[int]edge)
	}
	cg.edges[f][t] = edge{f, t, c}
}

type flow struct{ f, t, v int }

func findMaxFlow(g *capacityGraph) []flow {
	rg := newResidualGraph(g)
	return rg.findMaxFlow()
}

func (rg *residualGraph) findMaxFlow() []flow {
	for true { // find augmenting path
		fs, e := rg.findAugmentingPath()
		if e != nil {
			break
		}
		for _, f := range fs {
			rg.addFlowPath(f)
		}
	}
	return rg.listFlows()
}

type queue struct {
	rIdx, wIdx int
	queues     [2][][]int
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) isEmpty() bool {
	return q.rIdx == q.wIdx && len(q.queues[q.rIdx]) == 0
}

func (q *queue) push(p []int) {
	if q.rIdx == q.wIdx {
		q.wIdx = (q.wIdx + 1) % 2
	}
	q.queues[q.wIdx] = append(q.queues[q.wIdx], p)
}

func (q *queue) pop() (p []int) {
	if len(q.queues[q.rIdx]) == 0 {
		q.rIdx = q.wIdx
	}

	p, q.queues[q.rIdx] = q.queues[q.rIdx][0], q.queues[q.rIdx][1:len(q.queues[q.rIdx])]
	return
}

type residualGraph struct {
	*capacityGraph
	flows     map[int]map[int]flow
	residuals map[int]map[int]edge
}

func newResidualGraph(g *capacityGraph) *residualGraph {
	rg := &residualGraph{
		capacityGraph: g,
		flows:         make(map[int]map[int]flow),
		residuals:     make(map[int]map[int]edge),
	}
	for _, es := range g.edges {
		for _, e := range es {
			rg.setResidual(e.f, e.t, e.c)
		}
	}
	return rg
}

type marker struct{}

var mark = marker{}

func (rg *residualGraph) findAugmentingPath() ([]flow, error) {
	found := map[int]marker{}

	q := newQueue()
	q.push([]int{1})
	found[1] = mark
	for !q.isEmpty() {
		path := q.pop()
		current := path[len(path)-1]
		if current == rg.capacityGraph.nodeNum {
			return rg.flow(path), nil
		}
		for n, e := range rg.residuals[current] {
			if _, ok := found[n]; ok {
				continue
			}
			if e.c < 1 {
				continue
			}
			nextPath := make([]int, len(path))
			copy(nextPath, path)
			nextPath = append(nextPath, n)
			found[n] = mark
			q.push(nextPath)
		}
	}
	// bfs and min cap
	return nil, errors.New("not found")
}

func (rg *residualGraph) flow(crumb []int) (result []flow) {
	var min = math.MaxInt32
	for len(crumb) > 1 {
		f, t := crumb[0], crumb[1]
		res := rg.residuals[f][t]
		if res.c < min {
			min = res.c
		}
		result = append(result, flow{f, t, res.c})
		crumb = crumb[1:len(crumb)]
	}
	for i := range result {
		result[i].v = min
	}
	return
}

func (rg *residualGraph) listFlows() (result []flow) {
	for _, fs := range rg.flows {
		for _, f := range fs {
			result = append(result, f)
		}
	}
	return result
}

func (rg *residualGraph) addFlowPath(f flow) error {
	if rg.residuals[f.f][f.t].c < f.v {
		return errOverCapacity
	}
	rg.updateFlow(f)
	rg.updateResidual(f)
	return nil
}

var (
	errOverCapacity = errors.New("over capacity")
)

func (rg *residualGraph) updateFlow(f flow) {
	if _, ok := rg.flows[f.f]; !ok {
		rg.flows[f.f] = make(map[int]flow)
	}
	rg.flows[f.f][f.t] = flow{
		f: f.f,
		t: f.t,
		v: rg.flows[f.f][f.t].v + f.v,
	}
}

func (rg *residualGraph) updateResidual(f flow) {
	rg.setResidual(f.f, f.t, rg.residuals[f.f][f.t].c-f.v)
	rg.setResidual(f.t, f.f, rg.residuals[f.t][f.f].c+f.v)
}

func (rg *residualGraph) setResidual(f, t, c int) {
	if _, ok := rg.residuals[f]; !ok {
		rg.residuals[f] = make(map[int]edge)
	}
	rg.residuals[f][t] = edge{f, t, c}
}
