package main

import (
	"bytes"
	"fmt"
)

const (
	land  = '.'
	water = 'W'
)

func main() {
	var n, w int
	fmt.Scan(&n, &w)
	field := make([]string, n+2)
	visited := make([][]bool, n+2)

	skip := makeSafeField(w)
	field[0] = skip
	visited[0] = make([]bool, w+2)
	for i := 1; i <= n; i++ {
		var s string
		fmt.Scan(&s)
		field[i] = "." + s + "."
		visited[i] = make([]bool, w+2)
	}
	field[n+1] = skip
	visited[n+1] = make([]bool, w+2)

	num := 0
	// search start point
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if field[i][j] == 'W' && !visited[i][j] {
				num++
				dfs(i, j, field, visited)
			}
		}
	}
	fmt.Printf("%d\n", num)
}

func dfs(i, j int, field []string, visited [][]bool) {
	type d struct{ x, y int }
	visited[i][j] = true
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			ni, nj := i+di, j+dj
			if field[ni][nj] == 'W' && !visited[ni][nj] {
				dfs(ni, nj, field, visited)
			}
		}
	}
}

func makeSafeField(w int) string {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < w+2; i++ {
		buf.Write([]byte{'.'})
	}
	return buf.String()
}
