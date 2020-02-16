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
	p := parseProblem(os.Stdin)
	fmt.Printf("%d\n", p.findMinCost())
}

type problem struct {
	num          int
	cities       []city
	xInfo, yInfo *cityGroupInfos
}

func newProblem(cities []city) *problem {
	p := &problem{
		num:    len(cities),
		cities: cities,
		xInfo:  newCityGroupXInfos(cities),
		yInfo:  newCityGroupYInfos(cities),
	}
	return p
}

type cityGroupSummary map[int][]int

func newCityGroupXInfos(cities []city) *cityGroupInfos {
	summary := cityGroupSummary{}
	for id, city := range cities {
		summary[city.x] = append(summary[city.x], id)
	}
	infos := &cityGroupInfos{}
	for location, cities := range summary {
		infos.cities = append(infos.cities, cityGroupInfo{
			location: location,
			cities:   cities,
		})
	}
	sort.Sort(infos)
	return infos
}
func newCityGroupYInfos(cities []city) *cityGroupInfos {
	summary := cityGroupSummary{}
	for id, city := range cities {
		summary[city.y] = append(summary[city.y], id)
	}
	infos := &cityGroupInfos{}
	for location, cities := range summary {
		infos.cities = append(infos.cities, cityGroupInfo{
			location: location,
			cities:   cities,
		})
	}
	sort.Sort(infos)
	return infos
}

func (p *problem) startEdges(c int) (edges []edge) {
	from := p.cities[c]
	{
		ci := p.xInfo.cities[p.xInfo.Find(from.x)]
		for _, to := range ci.cities {
			if to == c {
				continue
			}
			cost := cost(from, p.cities[to])
			edges = append(edges, edge{cost: cost, from: c, to: to})
		}
	}
	{
		ci := p.yInfo.cities[p.yInfo.Find(from.y)]
		for _, to := range ci.cities {
			if to == c {
				continue
			}
			cost := cost(from, p.cities[to])
			edges = append(edges, edge{cost: cost, from: c, to: to})
		}
	}
	return edges
}

func (p *problem) candidate(c int) (edges []edge) {
	from := p.cities[c]
	for _, ci := range p.xInfo.nexts(from.x) {
		for _, to := range ci.cities {
			cost := cost(from, p.cities[to])
			edges = append(edges, edge{cost: cost, from: c, to: to})
		}
	}
	for _, ci := range p.yInfo.nexts(from.y) {
		for _, to := range ci.cities {
			cost := cost(from, p.cities[to])
			edges = append(edges, edge{cost: cost, from: c, to: to})
		}
	}
	return edges
}

var mark = struct{}{}

func (p *problem) findMinCost() int {
	// implement as Prim
	heap := NewHeap()
	id := 0
	for _, edge := range p.startEdges(id) {
		heap.Push(edge)
	}
	for _, edge := range p.candidate(id) {
		heap.Push(edge)
	}
	registered := map[int]struct{}{}
	registered[id] = mark
	treeCost := 0
	var addedEdges []edge
	for !heap.IsEmpty() && len(registered) < p.num {
		edge, _ := heap.Pop()
		// fmt.Printf("%d -> %d\n", edge.from, edge.to)
		if _, ok := registered[edge.to]; ok {
			// fmt.Println("SKIPEED")
			continue
		}
		treeCost += edge.cost
		addedEdges = append(addedEdges, edge)
		registered[edge.to] = mark
		for _, edge := range p.startEdges(edge.to) {
			if _, ok := registered[edge.to]; !ok {
				heap.Push(edge)
			}
		}
		for _, edge := range p.candidate(edge.to) {
			if _, ok := registered[edge.to]; !ok {
				heap.Push(edge)
			}
		}
		id = edge.to
	}
	return treeCost
}

type city struct {
	x, y int
}

func cost(from, to city) int {
	df := from.x - to.x
	costX := df ^ (df >> 31) - (df >> 31)
	dfy := from.y - to.y
	costY := dfy ^ (dfy >> 31) - (dfy >> 31)
	if costX <= costY {
		return costX
	}
	return costY
}

type edge struct {
	to, from, cost int
}

type cityGroupInfo struct {
	location int
	cities   []int
}

type cityGroupInfos struct {
	cities []cityGroupInfo
}

func (cs *cityGroupInfos) Len() int {
	return len(cs.cities)
}
func (cs *cityGroupInfos) Less(i, j int) bool {
	return cs.cities[i].location < cs.cities[j].location
}
func (cs *cityGroupInfos) Swap(i, j int) {
	cs.cities[i], cs.cities[j] = cs.cities[j], cs.cities[i]
}

func (cs *cityGroupInfos) Find(location int) (index int) {
	middle := len(cs.cities)/2 + 1
	return cs.find(location, middle, 1, len(cs.cities))
}

func (cs *cityGroupInfos) find(location, index, min, max int) (findIndex int) {
	if currentLocation := cs.cities[index-1].location; currentLocation == location {
		return index - 1
	} else {
		if currentLocation < location {
			next := (index+max)/2 + 1
			if index+1 > max {
				return cs.find(location, next, max, max)
			} else {
				return cs.find(location, next, index+1, max)
			}
		} else {
			next := (index + min) / 2
			if index-1 < min {
				return cs.find(location, next, min, min)
			} else {
				return cs.find(location, next, min, index-1)
			}
		}
	}
}

func (cs *cityGroupInfos) nexts(location int) (c []cityGroupInfo) {
	idx := cs.Find(location)
	for _, i := range []int{idx - 1, idx + 1} {
		if i >= 0 && i < len(cs.cities) {
			c = append(c, cs.cities[i])
		}
	}
	return
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
	return newProblem(cities)
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
