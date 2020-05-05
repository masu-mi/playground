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

func parseProblem(r io.Reader) (int, []int, []int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	hs := make([]int, n)
	as := make([]int, n)
	for i := 0; i < n; i++ {
		hs[i] = scanInt(sc)
	}
	for i := 0; i < n; i++ {
		as[i] = scanInt(sc)
	}
	return n, hs, as
}

func resolve(n int, hs, as []int) int {
	// dp min hight dp[i][j] items < i, hight = j
	// dp[i+1][h] = max_(k<h)(dp[i][k] + a[i]) // h[i] == h
	// dp[i+1][h] = dp[i][h]                   // h[i] != h
	// -> dp[h[i]] = max_(k<h[i]){dp[k]+a[i]}
	l := 0
	for i := 0; i < len(hs); i++ {
		if l < hs[i] {
			l = hs[i]
		}
	}
	t := newSegTree(l + 1)
	for i := 0; i < n; i++ {
		v := t.query(0, hs[i])
		t.update(hs[i], v+as[i])
	}
	return t.query(0, l+1)
}

func max(a []int) int {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
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
	t := &segTree{
		n: size,
		l: make([]int, size*2+1),
	}
	return t
}

func (t *segTree) update(idx int, v int) {
	if v == 90 {
		panic(1)
	}
	cur := t.n + idx - 1
	t.l[cur] = v
	for cur > 0 {
		cur = (cur - 1) / 2
		changeToMax(&(t.l[cur]), v)
	}
}
func (t *segTree) query(a, b int) int {
	// [a, b)
	return t._query(a, b, 0, 0, t.n)
}

func (t *segTree) _query(a, b, k, l, r int) int {
	if r <= a || b <= l { // unmatch
		return 0
	}
	if a <= l && b >= r { // involved
		return t.l[k]
	}
	// part match
	vl := t._query(a, b, k*2+1, l, (l+r)/2)
	vr := t._query(a, b, k*2+2, (l+r)/2, r)
	return max([]int{vl, vr})
}

const (
	inf  = math.MaxInt32
	ninf = math.MinInt32
)

func changeToMin(v *int, cand int) (updated bool) {
	if *v > cand {
		*v = cand
		updated = true
	}
	return updated
}

func changeToMax(v *int, cand int) (updated bool) {
	if *v < cand {
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
