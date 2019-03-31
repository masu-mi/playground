package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	grid := make([]string, h+2)
	wall := createWall(w)
	grid[0] = wall
	grid[h+1] = wall
	sc := bufio.NewScanner(os.Stdin)
	for i := 1; i <= h; i++ {
		sc.Scan()
		grid[i] = "." + sc.Text() + "."
	}
	for i := 1; i <= h; i++ {
		buf := bytes.NewBuffer([]byte{})
		for j := 1; j <= w; j++ {
			if grid[i][j] == '#' {
				buf.Write([]byte{'#'})
				continue
			}
			n := bombNum(grid, i, j)
			buf.WriteString(fmt.Sprintf("%d", n))
		}
		fmt.Println(buf.String())
	}
}

func bombNum(grid []string, i, j int) (num int) {
	type pos struct{ y, x int }
	for _, p := range []pos{
		{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1},
		{i, j - 1}, {i, j + 1},
		{i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1},
	} {
		if grid[p.y][p.x] == '#' {
			num++
		}
	}
	return
}

func createWall(w int) string {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < w+2; i++ {
		buf.Write([]byte{'.'})
	}
	return buf.String()
}
