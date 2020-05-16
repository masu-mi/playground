package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
			t.update(scanInt(sc), scanInt(sc))
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
	s := &segTree{n: size, l: make([]int, size*2-1)}
	for i := 0; i < len(s.l); i++ {
		s.l[i] = inf
	}
	return s
}

func (s *segTree) update(idx, v int) {
	cur := s.n + idx - 1
	s.l[cur] = v
	for cur > 0 {
		cur = (cur - 1) / 2
		changeToMin(&(s.l[cur]), v)
	}
}

func (s *segTree) query(a, b int) int {
	return s._query(a, b, 0, 0, s.n)
}

func (s *segTree) _query(a, b, k, l, r int) int {
	// [a, b), [l, r)
	if r <= a || b <= l { // unmatch
		return math.MaxInt32
	}
	if a <= l && r <= b { // direct match
		return s.l[k]
	}
	// sub match
	vl := s._query(a, b, k*2+1, l, (l+r)/2)
	vr := s._query(a, b, k*2+2, (l+r)/2, r)
	return min(vl, vr)
}

const (
	inf = math.MaxInt32
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func changeToMin(v *int, cand int) (updated bool) {
	if *v > cand {
		*v = cand
		updated = true
	}
	return updated
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
