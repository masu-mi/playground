package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s, g, grid := loadGrid(h, w)
	if bfsReachableWithBreak(h, w, 2, s, g, grid) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func bfsReachableWithBreak(h, w, limit int, s, g pos, grid []string) bool {
	visited := make([][]bool, h+2)
	for i := 0; i <= h+1; i++ {
		visited[i] = make([]bool, w+2)
	}

	q := newDeque(h * w)
	q.PushBack(candidate{s, -math.MaxInt32})
	for !q.Empty() {
		cur := q.Pop() // cur := q[0]
		visited[cur.p.y][cur.p.x] = true
		if eq(cur.p, g) {
			return cur.c+math.MaxInt32 <= limit
		}
		for _, d := range []pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := pos{cur.p.y + d.y, cur.p.x + d.x}
			if !(next.y >= 1 && next.y <= h && next.x >= 1 && next.x <= w) {
				continue
			}
			if visited[next.y][next.x] {
				continue
			}
			if grid[next.y][next.x] == '#' {
				if cur.c+1+math.MaxInt32 > limit {
					continue
				}
				q.PushBack(candidate{next, cur.c + 1})
			} else {
				if cur.c+math.MaxInt32 > limit {
					continue
				}
				q.PushFront(candidate{next, cur.c})
			}
		}
	}
	return false
}

type pos struct{ y, x int }

func eq(a, b pos) bool { return a.y == b.y && a.x == b.x }

type candidate struct {
	p pos
	c int
}

type deque struct {
	lq, rq []candidate
}

func newDeque(l int) *deque {
	return &deque{
		make([]candidate, 0, l),
		make([]candidate, 0, l),
	}
}

func (dq *deque) Empty() bool {
	return len(dq.lq)+len(dq.rq) == 0
}
func (dq *deque) PushBack(c candidate) {
	dq.rq = append(dq.rq, c)
}
func (dq *deque) PushFront(c candidate) {
	dq.lq = append(dq.lq, c)
}
func (dq *deque) Pop() (c candidate) {
	if len(dq.lq) > 0 {
		c = dq.lq[len(dq.lq)-1]
		dq.lq = dq.lq[0 : len(dq.lq)-1]
		return c
	}
	c = dq.rq[0]
	dq.rq = dq.rq[1:len(dq.rq)]
	return c
}

func wall(w int) string {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i <= w+1; i++ {
		buf.Write([]byte{'#'})
	}
	return buf.String()
}

func loadGrid(h, w int) (s, g pos, grid []string) {
	grid = make([]string, h+2)
	wl := wall(w)
	grid[0] = wl
	for i := 1; i <= h; i++ {
		var str string
		fmt.Scan(&str)
		grid[i] = "#" + str + "#"
		for j := 1; j <= w; j++ {
			if grid[i][j] == 's' {
				s = pos{i, j}
			} else if grid[i][j] == 'g' {
				g = pos{i, j}
			}
		}
	}
	grid[h+1] = wl
	return s, g, grid
}
