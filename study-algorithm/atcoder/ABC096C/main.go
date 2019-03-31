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
	// 0.0 0.1 .. 0.w, 0.w+1
	// 1.0 1.1 .. 1.w, 1.w+1
	// h.0 h.1 .. h.w, h.w+1
	// h+1.0 h+1.1 .. h+1.w, h+1.w+1
	// make wall
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < w+2; i++ {
		buf.Write([]byte{'.'})
	}
	grid[0] = buf.String()
	grid[h+1] = buf.String()
	sc := bufio.NewScanner(os.Stdin)
	for i := 1; i <= h; i++ {
		sc.Scan()
		grid[i] = "." + sc.Text() + "."
	}
	// check paintable
	paintable := true
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if grid[i][j] == '#' {
				if grid[i-1][j] == '.' &&
					grid[i][j-1] == '.' &&
					grid[i+1][j] == '.' &&
					grid[i][j+1] == '.' {
					paintable = false
				}
			}
		}
	}
	if paintable {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
