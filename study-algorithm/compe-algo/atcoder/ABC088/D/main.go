package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	m := loadMaze(h, w)
	fmt.Printf("%v\n", maxScore(m, pos{w - 1, h - 1}))
}

type pos struct {
	x, y int
}

func maxScore(maze []string, goal pos) int {
	sum := countBlock(maze)
	l, e := searchMinPathLength(maze, pos{0, 0}, goal)
	if e != nil {
		return -1
	}
	return (goal.x+1)*(goal.y+1) - sum - l
}

func countBlock(maze []string) (sum int) {
	for i := 0; i < len(maze); i++ {
		sum += strings.Count(maze[i], "#")
	}
	return sum
}

func searchMinPathLength(maze []string, start, goal pos) (int, error) {
	dists := make([][]int, goal.y+1)
	for y := 0; y < len(dists); y++ {
		dists[y] = make([]int, goal.x+1)
	}
	dfs(maze, pos{0, 0}, goal, dists)
	if l := dists[goal.y][goal.x]; l > 0 {
		return l, nil
	}
	return -1, errors.New("not found")
}

func dfs(maze []string, start, goal pos, dists [][]int) {
	dists[start.y][start.x] = 1

	var cur pos
	var candidates []pos

	candidates = append(candidates, start)
	for len(candidates) > 0 {
		cur, candidates = candidates[0], candidates[1:]
		if cur.x == goal.x && cur.y == goal.y {
			return
		}
		for _, diff := range []pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := pos{
				cur.x + diff.x,
				cur.y + diff.y,
			}
			if isBlocked(maze, next.x, next.y) || dists[next.y][next.x] > 0 {
				continue
			}
			dists[next.y][next.x] = dists[cur.y][cur.x] + 1
			candidates = append(candidates, next)
		}
	}
	return
}

func loadMaze(h, w int) []string {
	maze := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&(maze[i]))
	}
	return maze
}

func isBlocked(maze []string, x, y int) bool {
	if x < 0 || y < 0 || y >= len(maze) || x >= len(maze[y]) {
		return true
	}
	if maze[y][x] == '#' {
		return true
	}
	return false
}
