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
	p := parseProblem(os.Stdin)
	fmt.Printf("%d\n", p.findMinCost())
}

type problem struct {
	num    int
	cities []city
	edges  [][]edge
}

var mark = struct{}{}

func (p *problem) findMinCost() int {
	heap := NewHeap()
	city := 0
	for _, edge := range p.edges[city] {
		heap.Push(edge)
	}
	registered := map[int]struct{}{}
	registered[city] = mark
	treeCost := 0
	var addedEdges []edge
	for !heap.IsEmpty() && len(registered) < p.num {
		edge, _ := heap.Pop()
		if _, ok := registered[edge.to]; ok {
			continue
		}
		treeCost += edge.cost
		addedEdges = append(addedEdges, edge)
		registered[edge.to] = mark
		for _, edge := range p.edges[edge.to] {
			if _, ok := registered[edge.to]; !ok {
				heap.Push(edge)
			}
		}
	}
	return treeCost
}

type heap struct {
	num   int
	array []edge
}

func NewHeap() *heap {
	return &heap{}
}

func (h *heap) Push(e edge) {
	h.num++
	h.array = append(h.array, e)
	h.recoverRankWithPush()
}
func (h *heap) Pop() (edge, error) {
	if h.num == 0 {
		return edge{}, errors.New("empty")
	}
	head := h.array[0]
	h.array[0] = h.array[h.num-1]
	h.num--
	h.array = h.array[0:h.num]
	h.recoverRankWithPop()
	return head, nil
}

func (h *heap) IsEmpty() bool {
	return h.num == 0
}

func (h *heap) recoverRankWithPop() {
	location := 1
	scl := location * 2
	bcl := location*2 + 1
	for scl <= h.num {
		cIdx, sIdx, bIdx := location-1, scl-1, bcl-1
		cCost := h.array[cIdx].cost
		sCost := h.array[sIdx].cost
		if bIdx < h.num {
			bCost := h.array[bIdx].cost
			if cCost > sCost && cCost > bCost {
				if sCost <= bCost {
					h.array[cIdx], h.array[sIdx] = h.array[sIdx], h.array[cIdx]
					location = scl
				} else {
					h.array[cIdx], h.array[bIdx] = h.array[bIdx], h.array[cIdx]
					location = bcl
				}
			} else if cCost > sCost {
				h.array[cIdx], h.array[sIdx] = h.array[sIdx], h.array[cIdx]
				location = scl
			} else if cCost > bCost {
				h.array[cIdx], h.array[bIdx] = h.array[bIdx], h.array[cIdx]
				location = bcl
			} else {
				break
			}
		} else if cCost > sCost {
			h.array[cIdx], h.array[sIdx] = h.array[sIdx], h.array[cIdx]
			location = scl
		} else {
			break
		}
		scl = location * 2
		bcl = location*2 + 1
	}
}

func (h *heap) recoverRankWithPush() {
	location := h.num
	for location > 1 {
		pLocation := location / 2

		cIdx, pIdx := location-1, pLocation-1
		if h.array[pIdx].cost <= h.array[cIdx].cost {
			break
		}
		h.array[pIdx], h.array[cIdx] = h.array[cIdx], h.array[pIdx]
		location = pLocation
	}
}

func NewProblem(cities []city) *problem {
	p := &problem{
		num:    len(cities),
		cities: cities,
		edges:  make([][]edge, len(cities)),
	}
	for i, s := range cities {
		for j, t := range cities {
			if i == j {
				continue
			}
			p.edges[i] = append(
				p.edges[i],
				edge{to: j, from: i, cost: cost(s, t)},
			)
		}
	}
	return p
}

type city struct {
	x, y int
}

func cost(s, t city) int {
	xdiff := s.x - t.x
	xc := xdiff ^ (xdiff >> 31) - (xdiff >> 31)
	ydiff := s.y - t.y
	yc := ydiff ^ (ydiff >> 31) - (ydiff >> 31)
	if xc < yc {
		return xc
	}
	return yc
}

type edge struct {
	to, from, cost int
}

func parseProblem(r io.Reader) *problem {
	var n int
	fmt.Fscan(r, &n)
	cities := make([]city, 0, n)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		y, _ := strconv.Atoi(sc.Text())
		cities = append(cities, city{x: x, y: y})
	}
	return NewProblem(cities)
}
