package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)
	c, g := loadGrid(h, w, n)

	time := 0
	for i := 0; i <= n-1; i++ {
		start, goal := c[i], c[i+1]
		times := initTimes(h, w)
		time += bfs(start, goal, times, g)
	}
	fmt.Printf("%d\n", time)
}

func initTimes(h, w int) [][]int {
	ts := make([][]int, h+2)
	for i := 0; i < h+2; i++ {
		ts[i] = make([]int, w+2)
	}
	return ts
}

type pos struct{ y, x int }

func eq(a, b pos) bool {
	return a.y == b.y && a.x == b.x
}

func bfs(start, goal pos, times [][]int, grid []string) int {
	var q []pos
	times[start.y][start.x] = 0
	q = append(q, start)
	for len(q) > 0 {
		cur := q[0]
		if eq(goal, cur) {
			return times[cur.y][cur.x]
		}
		for _, move := range []pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			next := pos{cur.y + move.y, cur.x + move.x}
			if grid[next.y][next.x] != 'X' && times[next.y][next.x] == 0 {
				times[next.y][next.x] = times[cur.y][cur.x] + 1
				q = append(q, next)
			}
		}
		q = q[1:len(q)]
	}
	panic(-1)
	return 0
}

func loadGrid(h, w, n int) (checks []pos, grid []string) {

	checks = make([]pos, n+1)
	grid = make([]string, h+2)
	wl := wall(w)
	grid[0] = wl
	for i := 1; i <= h; i++ {
		var s string
		fmt.Scan(&s)
		grid[i] = "X" + s + "X"
		for j := 1; j <= w; j++ {
			if b := grid[i][j]; b != '.' && b != 'X' {
				p := pos{i, j}
				if b == 'S' {
					checks[0] = p
				} else {
					idx, _ := strconv.Atoi(string([]byte{b}))
					checks[idx] = p
				}
			}
		}
	}
	grid[h+1] = wl
	return checks, grid
}

func wall(w int) string {
	buf := bytes.NewBuffer(make([]byte, 0, w+2))
	for i := 0; i < w+2; i++ {
		buf.Write([]byte{'X'})
	}
	return buf.String()
}
