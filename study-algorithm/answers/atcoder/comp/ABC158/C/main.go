package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (a, b int) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	a = scanInt(sc)
	b = scanInt(sc)
	return
}

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func resolve(a, b int) int {
	maxPrice := max(originMax8(a), originMax10(b))
	for i := 1; i <= maxPrice; i++ {
		if a == (i*2)/25 && b == (i/10) {
			return i
		}
	}
	return -1
}

func originMax8(tax int) int {
	return ((tax + 1) * 100 / 8)
}
func originMax10(tax int) int {
	return ((tax + 1) * 10)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
