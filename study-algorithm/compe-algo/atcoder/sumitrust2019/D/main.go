package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (n int, s string) {
	fmt.Fscanf(r, "%d\n%s", &n, &s)
	return n, s
}

func resolve(n int, str string) int {
	var count int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				if find(i, j, k, str) {
					count++
				}
			}
		}
	}
	return count
}

func find(i, j, k int, str string) bool {
	target := []int{i, j, k}
	tIdx := 0
	for idx := 0; idx < len(str); idx++ {
		if int(str[idx]-'0') == target[tIdx] {
			tIdx++
		}
		if tIdx == len(target) {
			return true
		}
	}
	return false
}
