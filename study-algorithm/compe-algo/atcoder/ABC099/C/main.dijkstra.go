package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) int {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	return scanInt(sc)
}

func resolve(n int) int {
	states := make([]int, n+1)
	for i := 1; i <= n; i++ {
		states[i] = inf
	}
	states[0] = 0

	q := &candidates{}
	heap.Init(q)
	heap.Push(q, node{0, 0})
	for q.Len() > 0 {
		cur := heap.Pop(q).(node)
		if cur.cost > states[cur.id] {
			continue
		}
		for i := 1; cur.id+i <= n; i *= 6 {
			if changeToMin(&(states[cur.id+i]), cur.cost+1) {
				heap.Push(q, node{cur.id + i, cur.cost + 1})
			}
		}
		for i := 1; cur.id+i <= n; i *= 9 {
			if changeToMin(&(states[cur.id+i]), cur.cost+1) {
				heap.Push(q, node{cur.id + i, cur.cost + 1})
			}
		}
	}
	return states[n]
}

const (
	inf = math.MaxInt32
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

type node struct {
	id, cost int
}

type candidates []node

func (h *candidates) Len() int {
	return len(*h)
}

func (h *candidates) Less(i, j int) bool {
	items := *h
	return items[i].cost < items[j].cost
}

func (h *candidates) Swap(i, j int) {
	items := *h
	items[i], items[j] = items[j], items[i]
}

func (h *candidates) Push(x interface{}) {
	*h = append(*h, x.(node))
}

func (h *candidates) Pop() interface{} {
	items := *h
	l := items[len(items)-1]
	*h = items[0 : len(items)-1]
	return l
}

func (h *candidates) top() node {
	return (*h)[h.Len()-1]
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
