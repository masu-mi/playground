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
	nums := make([][]int, h+2)
	for i := 0; i <= h+1; i++ {
		nums[i] = make([]int, w+2)
	}
	broken := bfsMinNumOfBrokenWall(h, w, s, g, nums, grid)
	if broken <= 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

type pos struct{ y, x int }

func eq(a, b pos) bool { return a.y == b.y && a.x == b.x }

func bfsMinNumOfBrokenWall(h, w int, s, g pos, brokenNums [][]int, grid []string) int {
	q := []pos{s}
	brokenNums[s.y][s.x] = math.MinInt32
	for len(q) > 0 {
		cur := q[0]
		for _, d := range []pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := pos{cur.y + d.y, cur.x + d.x}
			if !(next.y >= 1 && next.y <= h && next.x >= 1 && next.x <= w) {
				continue
			}
			curScore := brokenNums[cur.y][cur.x]
			if grid[next.y][next.x] == '#' {
				curScore++
			}
			if brokenNums[next.y][next.x] > curScore {
				brokenNums[next.y][next.x] = curScore
				q = append(q, next)
			}
		}
		q = q[1:len(q)]
	}
	return brokenNums[g.y][g.x] + math.MaxInt32 + 1
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
