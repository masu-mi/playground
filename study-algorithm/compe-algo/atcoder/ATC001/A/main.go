package main

import (
	"bytes"
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	field := make([]string, h+2)
	visited := make([][]bool, h+2)

	wall := makeWall(w)
	field[0] = wall
	visited[0] = make([]bool, w+2)
	for i := 1; i <= h; i++ {
		var s string
		fmt.Scan(&s)
		field[i] = "#" + s + "#"
		visited[i] = make([]bool, w+2)
	}
	field[h+1] = wall
	visited[h+1] = make([]bool, w+2)

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if field[i][j] == start {
				if dfs(i, j, field, visited) {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
	return
}

func dfs(i, j int, field []string, visited [][]bool) bool {
	visited[i][j] = true
	type move struct{ v, h int }
	for _, m := range []move{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		ni, nj := i+m.v, j+m.h
		if field[ni][nj] == goal {
			return true
		} else if field[ni][nj] != wall && !visited[ni][nj] {
			if dfs(ni, nj, field, visited) {
				return true
			}
		}
	}
	return false
}

const (
	start = 's'
	goal  = 'g'
	wall  = '#'
)

func makeWall(w int) string {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i <= w+1; i++ {
		buf.Write([]byte{wall})
	}
	return buf.String()
}
