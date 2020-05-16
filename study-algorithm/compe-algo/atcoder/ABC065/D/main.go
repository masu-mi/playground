package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	r := os.Stdin
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	var n int
	fmt.Fscanf(r, "%d", &n)
	var x, y int
	for i := 0; i < n; i++ {

		sc.Scan()
		x, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		y, _ = strconv.Atoi(sc.Text())

		cities = append(cities, city{x, y})
	}
	xList := make([]int, len(cities))
	for i := 0; i < len(cities); i++ {
		xList[i] = i
	}
	yList := make([]int, len(cities))
	for i := 0; i < len(cities); i++ {
		yList[i] = i
	}
	candX := cityList{xList, func(id int) int {
		return cities[id].x
	}}.edges()
	candY := cityList{yList, func(id int) int {
		return cities[id].y
	}}.edges()
	_, cost, _ := findMSTWithKruskal(len(cities), append(candX, candY...))
	fmt.Printf("%d\n", cost)
}

type city struct{ x, y int }

var cities []city

func cost(s, t city) int {
	costX := abs(s.x - t.x)
	costY := abs(s.y - t.y)
	if costX < costY {
		return costX
	}
	return costY
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

type cityList struct {
	list []int
	comp func(id int) int
}

func (cl cityList) edges() (edges []edge) {
	sort.Sort(cl)
	for i := 0; i < len(cl.list)-1; i++ {
		x, y := cl.list[i], cl.list[i+1]
		edges = append(edges, edge{x, y, cost(cities[x], cities[y])})
	}
	return
}

func (cl cityList) Len() int {
	return len(cl.list)
}

func (cl cityList) Less(i, j int) bool {
	return cl.comp(cl.list[i]) < cl.comp(cl.list[j])
}

func (cl cityList) Swap(i, j int) {
	cl.list[i], cl.list[j] = cl.list[j], cl.list[i]
}

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

func newUnifonFind(card int) *unionfind {
	uf := &unionfind{
		card:      card,
		parent:    make([]int, card),
		rank:      make([]int, card),
		childs:    make([]int, card),
		connected: card == 1,
	}
	for i := 0; i < card; i++ {
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
