package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	h, w, grid := loadGrid(os.Stdin)

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

func loadGrid(r io.Reader) (h, w int, grid []string) {
	fmt.Fscan(r, &h, &w)
	grid = make([]string, h+2)
	wall := createWall(w)
	grid[0] = wall
	sc := bufio.NewScanner(r)
	for i := 1; i <= h; i++ {
		sc.Scan()
		buf := bytes.NewBuffer([]byte{})
		buf.Write([]byte{'.'})
		buf.WriteString(sc.Text())
		buf.Write([]byte{'.'})
		grid[i] = buf.String()
	}
	grid[h+1] = wall
	return h, w, grid
}
func createWall(w int) string {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < w+2; i++ {
		buf.Write([]byte{'.'})
	}
	return buf.String()
}
