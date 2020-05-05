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
	ps := make([]int, m)
	ys := make([]int, m)
	pys := make([][]int, n)
	for i := 0; i < m; i++ {
		ps[i] = scanInt(sc)
		ps[i]-- // 0-indexed
		ys[i] = scanInt(sc)
		pys[ps[i]] = append(pys[ps[i]], ys[i])
	}
	for i := 0; i < n; i++ {
		sort.Sort(sort.IntSlice(pys[i]))
		uniq(pys[i])
	}
	for i := 0; i < m; i++ {
		// reverse: 1-indexed from 0-indexed
		fmt.Printf("%06d", ps[i]+1)
		idx := sort.Search(len(pys[ps[i]]), func(idx int) bool {
			return pys[ps[i]][idx] >= ys[i]
		})
		fmt.Printf("%06d\n", idx+1)
	}
	return
}

func uniq(l []int) {
	if len(l) == 0 {
		return
	}
	r := make([]int, 0, len(l))
	pre := l[0]
	r = append(r, l[0])
	for i := 1; i < len(l); i++ {
		if pre == l[i] {
			continue
		}
		r = append(r, l[i])
	}
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
