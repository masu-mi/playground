package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	resolve(os.Stdin)
}

func resolve(r io.Reader) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, m := scanInt(sc), scanInt(sc)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = scanInt(sc)
	}
	if !sort.IsSorted(sort.IntSlice(nums)) {
		fmt.Printf("Invalid Input %v\n", nums)
		os.Exit(1)
	}
	i := lowerBound(len(nums), func(i int) bool {
		return nums[i] >= m
	})
	fmt.Printf(">= m: %d at %d, %v\n", m, i, nums)

	i = upperBound(len(nums), func(i int) bool {
		return nums[i] < m
	})
	fmt.Printf(" < m: %d at %d, %v\n", m, i, nums)

	i = upperBound(len(nums), func(i int) bool {
		return nums[i] <= m
	})
	fmt.Printf("<= m: %d at %d, %v\n", m, i, nums)
}

func lowerBound(n int, f func(i int) bool) int {
	return sort.Search(n, f)
}

func upperBound(n int, f func(i int) bool) int {
	return sort.Search(n, func(i int) bool { return !f(i) }) - 1
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
