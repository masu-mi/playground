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

func parseProblem(r io.Reader) (n int, a, b []int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	n = scanInt(sc)
	return n, scanIntSlice(sc, n), scanIntSlice(sc, n)
}

func scanIntSlice(sc *bufio.Scanner, n int) (s []int) {
	for i := 0; i < n; i++ {
		s = append(s, scanInt(sc)-1)
	}
	return s
}

func resolve(n int, a, b []int) int {
	numA, numB := -1, -1
	i := 0
	for l := range permutations(n) {
		if match(l, a) {
			numA = i
		}
		if match(l, b) {
			numB = i
		}
		if numA > -1 && numB > -1 {
			break
		}
		i++
	}
	return abs(numA - numB)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func match(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func permutations(l int) chan []int {
	ch := make(chan []int)
	go func() {
		dfsPermutations(0, make([]bool, l), []int{}, func(perm []int) bool {
			p := make([]int, len(perm))
			copy(p, perm)
			ch <- p
			return false
		})
		close(ch)
	}()
	return ch
}

func dfsPermutations(pos int, used []bool, perm []int, atLeaf func(perm []int) (halt bool)) (halt bool) {
	l := len(used)
	if pos == l {
		if atLeaf(perm) {
			return true
		}
	}

	for i := 0; i < l; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		if dfsPermutations(pos+1, used, append(perm, i), atLeaf) {
			return true
		}
		used[i] = false
	}
	return false
}
