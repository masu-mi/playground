package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (a, b, ab, x, y int) {
	fmt.Fscanf(r, "%d %d %d %d %d", &a, &b, &ab, &x, &y)
	return
}

func resolve(a, b, ab, x, y int) int {
	return min([]int{
		a*x + b*y,
		2*ab*x + b*reLU(y-x),
		2*ab*y + a*reLU(x-y),
		2 * ab * max(x, y),
	})
}

func reLU(a int) int {
	if a < 0 {
		return 0
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a []int) int {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}
