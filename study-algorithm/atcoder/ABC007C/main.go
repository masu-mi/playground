package main

import (
	"bytes"
	"errors"
	"fmt"
)

type pos struct{ y, x int }

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	start := pos{}
	fmt.Scan(&(start.y), &(start.x))
	goal := pos{}
	fmt.Scan(&goal.y, &goal.x)
	grid := loadGrid(r, c)
	fmt.Printf("%d\n", bfs(start, goal, r, c, grid))
}

func bfs(s, g pos, r, c int, field []string) int {
	dists := make([][]int, r+2)
	for i := 0; i <= r+1; i++ {
		dists[i] = make([]int, c+2)
	}
	q := NewQueue()
	q.Push(s)
	for !q.Empty() {
		cur, _ := q.Pop()
		if cur.y == g.y && cur.x == g.x {
			return dists[cur.y][cur.x]
		}
		for _, d := range []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			next := pos{cur.y + d.y, cur.x + d.x}
			if field[next.y][next.x] == '.' && dists[next.y][next.x] == 0 && (next.y != s.y || next.x != s.x) {
				dists[next.y][next.x] = dists[cur.y][cur.x] + 1
				q.Push(next)
			}
		}
	}
	return -1
}

type queue struct {
	rIdx, wIdx int
	qs         [][]pos
}

func NewQueue() *queue {
	q := &queue{}
	q.qs = make([][]pos, 2)
	return q
}

func (q *queue) Push(p pos) {
	if q.rIdx == q.wIdx {
		q.wIdx = (q.wIdx + 1) % 2
	}
	q.qs[q.wIdx] = append(q.qs[q.wIdx], p)
}

func (q *queue) Empty() bool {
	return q.rIdx == q.wIdx && len(q.qs[q.rIdx]) == 0
}

func (q *queue) Pop() (p pos, e error) {
	if q.rIdx != q.wIdx && len(q.qs[q.rIdx]) == 0 {
		q.qs[q.rIdx] = nil
		q.rIdx = q.wIdx
	}
	if len(q.qs[q.rIdx]) == 0 {
		return pos{}, errors.New("empty")
	}
	p, q.qs[q.rIdx] = q.qs[q.rIdx][0], q.qs[q.rIdx][1:len(q.qs[q.rIdx])]
	return p, nil
}

func loadGrid(r, c int) []string {
	grid := make([]string, r+2)
	w := wall(c)
	grid[0] = w
	grid[r+1] = w
	for i := 1; i <= r; i++ {
		var s string
		fmt.Scan(&s)
		grid[i] = "#" + s + "#"
	}
	return grid
}

func wall(c int) string {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < c+2; i++ {
		buf.Write([]byte{'#'})
	}
	return buf.String()
}
