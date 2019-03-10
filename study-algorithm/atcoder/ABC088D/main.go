package main

import (
	"bytes"
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

type history struct {
	pos
	l int
}

func searchMinPathLength(maze []string, goal pos) (int, error) {
	// return min path, if not found this returns -1, errors.New("not found")
	cur := pos{0, 0}
	candidates := []history{}
	curHist := history{}
	note(maze, cur)
	candidates = append(candidates, history{cur, 1})
	for len(candidates) > 0 {
		curHist = candidates[0]
		candidates = candidates[1:]
		if curHist.x == goal.x && curHist.y == goal.y {
			return curHist.l, nil
		}
		for _, n := range nextAvailables(maze, curHist.pos) {
			note(maze, n)
			candidates = append(candidates, history{n, curHist.l + 1})
		}
	}
	return -1, errors.New("not found")
}

func note(maze []string, cur pos) {
	var buf bytes.Buffer
	l := maze[cur.y]
	buf.WriteString(l[0:cur.x])
	buf.WriteString("#")
	buf.WriteString(l[cur.x+1 : len(l)])
	maze[cur.y] = buf.String()
}

func nextAvailables(maze []string, cur pos) (candidates []pos) {
	for _, diff := range []pos{pos{0, 1}, pos{0, -1}, pos{1, 0}, pos{-1, 0}} {
		if isBlocked(maze, cur.x+diff.x, cur.y+diff.y) {
			continue
		}
		candidates = append(candidates, pos{cur.x + diff.x, cur.y + diff.y})
	}
	return candidates
}

func maxScore(maze []string, goal pos) int {
	sum := countBlock(maze)
	l, e := searchMinPathLength(maze, goal)
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
