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
	n := scanInt(sc)
	t := newSegTree(n, math.MaxInt32)
	for i := 0; i < n; i++ {
		t.update(i, scanInt(sc))
	}
	m := scanInt(sc)
	for i := 0; i < m; i++ {
		fmt.Println(t.query(scanInt(sc), scanInt(sc)))
	}
	return
}

func resolve(n int) int {
	return n
}

type segTree []int

func newSegTree(n int, init int) *segTree {
	// 0-indexed
	var l segTree = make([]int, 2*n+1)
	for i := 0; i < 2*n+1; i++ {
		l[i] = init
	}
	return &l
}

func (t *segTree) query(a, b int) int {
	return t._query(a, b, 0, 0, (len(*t)-1)/2)
}

func (t *segTree) _query(a, b, k int, l, r int) int {
	if k < 0 {
		// assert
		panic(1)
	}
	if r <= a || b <= l {
		return t.defautl()
	}
	if a <= l && r <= b {
		return (*t)[k]
	}
	vl := t._query(a, b, k*2+1, l, (l+r)/2)
	vr := t._query(a, b, k*2+2, (l+r)/2, r)
	return t._judge(vl, vr)
}

func (t *segTree) update(idx int, v int) {
	list := []int(*t)
	cur := (len(list)-1)/2 + idx - 1
	list[cur] = v
	for cur > 0 {
		cur = (cur - 1) / 2
		t._update(cur, v)
	}
}

func (t *segTree) defautl() int {
	return math.MaxInt32
}

func (t *segTree) _judge(vl, vr int) int {
	if vl < vr {
		return vl
	}
	return vr
}

func (t *segTree) _update(idx, v int) {
	list := []int(*t)
	if list[idx] > v {
		list[idx] = v
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
