package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/k0kubun/pp"
)

func main() {
	fmt.Println("vim-go")
	p := parseProblem(os.Stdin)
	fmt.Println("[ Problem ]")
	pp.Println(p)
	r, _ := searchStableMatching(*p)
	fmt.Println("[ Result ]")
	pp.Println(r)
}

type edge struct{ f, s int }

type marker struct{}

var mark = marker{}

func searchStableMatching(p problem) ([]edge, error) {
	g := newGaleShapley(p)
	return g.search()
}

type galeShapley struct {
	keeped   map[id]id
	rejected map[id]map[id]marker
	p        problem
}

func newGaleShapley(p problem) *galeShapley {
	return &galeShapley{
		keeped:   make(map[id]id),
		rejected: make(map[id]map[id]marker),
		p:        p,
	}
}

func (g *galeShapley) search() ([]edge, error) {
	q := newQueue()
	for i := 0; i < g.p.num; i++ {
		q.push(id{1, i + 1})
	}
	for !q.empty() {
		f, _ := q.pop()
		if _, ok := g.keeped[f]; ok {
			return nil, errors.New("internal: keepded candidate exists in queue")
		}
		next, e := g.nextCandidate(f)
		if e != nil {
			fmt.Fprintf(os.Stderr, e.Error())
			os.Exit(1)
		}
		nextPair, ok := g.keeped[next]
		if !ok {
			g.keeped[next], g.keeped[f] = f, next
			continue
		}
		if g.p.orders[next][nextPair.num] > g.p.orders[next][f.num] {
			g.rejected[nextPair][next] = mark
			g.keeped[next], g.keeped[f] = f, next
			delete(g.keeped, nextPair)
			q.push(nextPair)
		} else {
			g.rejected[f][next] = mark
			q.push(f)
		}
	}
	var result []edge
	for k, v := range g.keeped {
		if k.cat == 2 {
			continue
		}
		result = append(result, edge{f: k.num, s: v.num})
	}
	return result, nil
}

func (g *galeShapley) nextCandidate(fID id) (id, error) {
	for _, sName := range g.p.candidates[fID] {
		sID := id{2, sName}
		if g.isRejected(fID, sID) {
			continue
		}
		return sID, nil
	}
	return id{0, 0}, errors.New("no candidate")
}

func (g *galeShapley) isRejected(fID, sID id) bool {
	if _, ok := g.rejected[fID]; !ok {
		g.rejected[fID] = map[id]marker{}
	}
	_, ok := g.rejected[fID][sID]
	return ok
}

type id struct{ cat, num int }

type problem struct {
	num int

	candidates map[id][]int
	orders     map[id]map[int]int
}

func newProblem() *problem {
	return &problem{
		candidates: make(map[id][]int),
		orders:     make(map[id]map[int]int),
	}
}

func parseProblem(r io.Reader) *problem {
	p := newProblem()
	fmt.Fscan(r, &p.num)
	for i := 0; i < p.num; i++ {
		c, o := parseOrder(r, p.num)
		fID := id{1, i + 1}
		p.candidates[fID] = c
		p.orders[fID] = o
	}
	for i := 0; i < p.num; i++ {
		c, o := parseOrder(r, p.num)
		sID := id{2, i + 1}
		p.candidates[sID] = c
		p.orders[sID] = o
	}
	return p
}

func parseOrder(r io.Reader, n int) ([]int, map[int]int) {
	candidate := []int{}
	order := map[int]int{}
	for i := 0; i < n; i++ {
		var cand int
		fmt.Fscan(r, &cand)

		candidate = append(candidate, cand)
		order[cand] = i + 1
	}
	return candidate, order
}

type queue struct {
	rIdx, wIdx int
	qs         [][]id
}

func newQueue() *queue {
	return &queue{qs: make([][]id, 2)}
}

func (q *queue) push(item id) {
	if q.rIdx == q.wIdx {
		q.wIdx = (q.wIdx + 1) % 2
	}
	q.qs[q.wIdx] = append(q.qs[q.wIdx], item)
}
func (q *queue) empty() bool {
	return q.rIdx == q.wIdx && len(q.qs[q.rIdx]) == 0
}
func (q *queue) pop() (item id, e error) {
	if q.rIdx != q.wIdx && len(q.qs[q.rIdx]) == 0 {
		q.qs[q.rIdx] = nil
		q.rIdx = q.wIdx
	}
	if len(q.qs[q.rIdx]) == 0 {
		return id{0, 0}, errors.New("empty")
	}
	item, q.qs[q.rIdx] = q.qs[q.rIdx][0], q.qs[q.rIdx][1:len(q.qs[q.rIdx])]
	return item, nil
}
