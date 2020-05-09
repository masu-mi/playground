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
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (int, map[int][]query) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, m := scanInt(sc), scanInt(sc)
	qs := make(map[int][]query, m)
	for i := 0; i < m; i++ {
		l, r, a := scanInt(sc), scanInt(sc)+1, scanInt(sc)
		qs[r] = append(qs[r], query{l: l, r: r, a: a})
	}
	return n, qs
}

type query struct {
	l, r, a int
}

func resolve(n int, qs map[int][]query) int {
	t := newLazySegTree(n + 1)
	for i := 1; i <= n+1; i++ {
		max := t.max(0, i)
		sum := 0
		for _, v := range qs[i] {
			sum += v.a
		}
		t.update(i, i+1, max+sum-t.max(i, i+1))
		for _, v := range qs[i] {
			t.update(v.l, i, v.a)
		}
	}
	return t.max(0, n+2)
}

type lazySegTree struct {
	size int
	act  []int
	mon  []int
}

func newLazySegTree(last int) *lazySegTree {
	num := last + 1
	size := 1
	for size < num {
		size <<= 1
	}
	return &lazySegTree{
		size: size,
		act:  make([]int, size*2-1),
		mon:  make([]int, size*2-1),
	}
}

func (t *lazySegTree) _idx(idx int) int {
	return idx + t.size - 1
}

func (t *lazySegTree) max(a, b int) int {
	return t._max(a, b, 0, 0, t.size)
}
func (t *lazySegTree) _max(a, b, k, l, r int) int {
	t.evaluate(k)
	if a <= l && r <= b {
		return t.mon[k]
	} else if a < r && l < b {
		return max(t._max(a, b, 2*k+1, l, (l+r)>>1), t._max(a, b, 2*k+2, (l+r)>>1, r))
	}
	return math.MinInt32
}

func (t *lazySegTree) update(a, b, v int) {
	t._update(a, b, v, 0, 0, t.size)
}
func (t *lazySegTree) _update(a, b, v, k, l, r int) {
	t.evaluate(k)
	if a <= l && r <= b {
		t.act[k] += v
		t.evaluate(k)
	} else if a < r && l < b {
		t._update(a, b, v, 2*k+1, l, (l+r)/2)
		t._update(a, b, v, 2*k+2, (l+r)/2, r)
		t.mon[k] = max(t.mon[2*k+1], t.mon[2*k+2])
	}
}

func (t *lazySegTree) evaluate(k int) {
	if t.act[k] == 0 {
		return
	}
	if k < t.size-1 {
		t.act[2*k+1] += t.act[k]
		t.act[2*k+2] += t.act[k]
	}
	t.mon[k] += t.act[k]
	t.act[k] = 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
