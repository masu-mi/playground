package kruskal

import (
	"errors"
	"sort"
)

func findMSTWithKruskal(card int, edges edgeList) (mst edgeList, cost int, err error) {
	sort.Sort(edges)
	l := len(edges)
	uf := newUnifonFind(card)
	for i := 0; !uf.connected && i < l; i++ {
		e := edges[i]
		if uf.same(e.x, e.y) {
			continue
		}
		uf.union(edges[i].x, edges[i].y)
		cost += e.cost
		mst = append(mst, e)
	}
	if !uf.connected {
		return mst, cost, errors.New("can't construct MST")
	}
	return mst, cost, nil
}

type edge struct {
	x, y, cost int
}

type edgeList []edge

func (e edgeList) Len() int {
	return len(e)
}

func (e edgeList) Less(i, j int) bool {
	return e[i].cost < e[j].cost
}

func (e edgeList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type unionfind struct {
	card      int
	connected bool

	parent []int
	rank   []int
	childs []int
}

func newUnifonFind(size int) *unionfind {
	uf := &unionfind{
		card:      size,
		parent:    make([]int, size),
		rank:      make([]int, size),
		childs:    make([]int, size),
		connected: size == 1,
	}
	for i := 0; i < size; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (u *unionfind) find(x int) int {
	p := u.parent[x]
	if p == x {
		return x
	}
	r := u.find(p)
	u.parent[x] = r
	return r
}

func (u *unionfind) same(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *unionfind) union(x, y int) {
	xR, yR := u.find(x), u.find(y)
	if xR == yR {
		return
	}
	if rankX, rankY := u.rank[xR], u.rank[yR]; rankX < rankY {
		u.parent[xR] = yR
		u.childs[yR] += u.childs[xR] + 1
		u.connected = u.card == u.childs[yR]+1
	} else {
		u.parent[yR] = xR
		u.childs[xR] += u.childs[yR] + 1
		u.connected = u.card == u.childs[xR]+1
		if rankX == rankY {
			u.rank[xR]++
		}
	}
}

func (u *unionfind) size(x int) int {
	return u.childs[u.find(x)] + 1
}
