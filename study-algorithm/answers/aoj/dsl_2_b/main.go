package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	parseProblem(os.Stdin)
}

func parseProblem(r io.Reader) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, q := scanInt(sc), scanInt(sc)
	t := newSegTree(n)
	for i := 0; i < q; i++ {
		op := scanInt(sc)
		if op == 0 {
			t.add(scanInt(sc), scanInt(sc))
		} else {
			fmt.Println(t.query(scanInt(sc), scanInt(sc)+1))
		}
	}
}

type segTree struct {
	n int
	l []int
}

func newSegTree(n int) *segTree {
	size := 1
	for size < n {
		size <<= 1
	}
	return &segTree{
		n: size,
		l: make([]int, 2*size-1),
	}
}

func (t *segTree) add(idx int, v int) {
	cur := idx + t.n - 1
	t.l[cur] += v
	for cur > 0 {
		cur = (cur - 1) / 2
		t.l[cur] = t.l[cur*2+1] + t.l[cur+2+2]
	}
}

func (t *segTree) query(a, b int) int {
	return t._query(a, b, 0, 0, t.n)
}

func (t *segTree) _query(a, b, k, l, r int) int {
	// [a, b), [l, r)
	if r <= a || b <= l { // unmatch
		return 0
	}
	if a <= l && r <= b { // involved
		return t.l[k]
	}
	vl := t._query(a, b, 2*k+1, l, (l+r)/2)
	vr := t._query(a, b, 2*k+2, (l+r)/2, r)
	return vl + vr
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
